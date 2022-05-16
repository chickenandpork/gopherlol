package commands

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	RegisterCommands(&TestMockExtensionCommands{})
}

// TestHandler codifies the expected behavior of the Handler() to ensure that the freedom to
// change/refactor/extend does not break current functionality
func TestHandler(t *testing.T) {
	assert.Nil(t, SetLogSink(t))
	tests := []struct {
		name     string
		query    string
		retcode  int
		expected string
	}{
		{
			"bok",
			"/?q=bokbokbok", http.StatusSeeOther,
			"<a href=\"https://www.google.com/#q=bokbokbok\">See Other</a>."},
		{
			"G20",
			"/?q=g%20bok", http.StatusSeeOther,
			"<a href=\"https://www.google.com/#q=bok\">See Other</a>."},
		{
			"Gplus",
			"/?q=g+bok", http.StatusSeeOther,
			"<a href=\"https://www.google.com/#q=bok\">See Other</a>."},
		{
			"Gspace",
			"/?q=g bok", http.StatusSeeOther,
			"<a href=\"https://www.google.com/#q=bok\">See Other</a>."},
		{
			"So20",
			"/?q=so%20awesome", http.StatusSeeOther,
			"<a href=\"https://stackoverflow.com/search?q=awesome\">See Other</a>."},
		{
			"SoPlus",
			"/?q=so+awesome", http.StatusSeeOther,
			"<a href=\"https://stackoverflow.com/search?q=awesome\">See Other</a>."},
		{
			"SoSpace",
			"/?q=so awesome", http.StatusSeeOther,
			"<a href=\"https://stackoverflow.com/search?q=awesome\">See Other</a>."},
		{
			"fallback",
			"/?q=mega+awesome", http.StatusSeeOther,
			"<a href=\"https://www.google.com/#q=mega+awesome\">See Other</a>."},
		{
			"help",
			"/?q=help", http.StatusOK,
			"<h1>gopherlol command list</h1><ul>" +
				"<li><strong>author</strong></li>" +
				"<li><strong>g</strong>, takes args</li>" +
				"<li><strong>help</strong></li>" +
				"<li><strong>list</strong></li>" +
				"<li><strong>so</strong>, takes args</li>" +
				"<li><strong>author</strong></li>" + // as noted, TODO: make only one author key
				"<li><strong>fbgs</strong>, takes args</li>" +
				"<li><strong>fullbiggrepsearch</strong>, takes args</li>" +
				"<li><strong>go</strong>, takes args</li>" +
				"</ul>"},
		{
			"list",
			"/?q=help", http.StatusOK,
			"<h1>gopherlol command list</h1><ul>" +
				"<li><strong>author</strong></li>" +
				"<li><strong>g</strong>, takes args</li>" +
				"<li><strong>help</strong></li>" +
				"<li><strong>list</strong></li>" +
				"<li><strong>so</strong>, takes args</li>" +
				"<li><strong>author</strong></li>" + // as noted, TODO: make only one author key
				"<li><strong>fbgs</strong>, takes args</li>" +
				"<li><strong>fullbiggrepsearch</strong>, takes args</li>" +
				"<li><strong>go</strong>, takes args</li>" +
				"</ul>"},
		{
			"Go",
			"/?q=go Errorf", http.StatusSeeOther,
			"<a href=\"https://go.example.com/Errorf\">See Other</a>."},
		{
			"regex",
			"/?q=ABC-1234", http.StatusSeeOther,
			"<a href=\"https://two.example.com/ABC-1234\">See Other</a>."},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Generate a fake request to pass to the handler under test.  I'm not sure
			// if the query parameters should be passed in rather than "nil", but
			// because the handler is expected to be called from a browser's search
			// util, it's probably fairly accurate to give a query string like a basic
			// GET or SEARCH method.
			req, err := http.NewRequest("GET", test.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ReponseRecorder via New~ to collect he response in the handler.
			// ResponseRecorder can cast to http.ResponseWriter so we can use it here
			//  rather than more complex scaffolding.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			// We can get away with directly calling the ServeHTTP(...) method, and have
			// it pass to the handler routine for servicing, because the Handler
			// satisfies http.Handler.  Our query string above ("/..." from
			// tests[].query) cause handling to be passed to commands.Handler()
			handler.ServeHTTP(rr, req)

			// Confirm the expected Status code was returned (it's all 302 redirects
			// right now, but that can change)
			assert.Equalf(t, rr.Code, test.retcode, "handler returned wrong status code: expected %v observed %v", test.retcode, rr.Code)

			// Check the response body is what we expect modulo some whitespace
			expected := strings.TrimSpace(test.expected)
			observed := strings.TrimSpace(rr.Body.String())
			assert.Equalf(t, expected, observed, "handler %s returned incorrect text: expected %v observed %v", test.name, expected, observed)
		})
	}
}

// TestMockExtensionCommands shows a possible extnesion of the commands and is intended to allow
// testing of modular extensions in general.  To see how little is needed to extend gopherlol, see
// github.com/chickenandpork/gopherlol-extend.  Copy -- not fork -- that repos and change as
// needed.
//
// For maintenance, I would of course suggest not copying the github.com/chickenandpork/gopherlol,
// just import is as gopherlol-extend does.  This allows you to update versions using an eventual
// go-compatible dependabot or similar to take advantage of any later improvements to this project
// with the least possible effort.
//
// Remember that we need exported symbols, which are capitalized in Go
type TestMockExtensionCommands struct{}

// Author permits sharing attribution or a URL for more info, or help
func (c *TestMockExtensionCommands) Author() string {
	return "https://github.com/chickenandpork"
}

// TryRegex looks for a XXX-1234 -style pattern
func (c *TestMockExtensionCommands) TryRegex(parm string) string {
	if m.Match([]byte(parm)) {
		return fmt.Sprintf("https://try.example.com/%s", parm)
	}
	return ""
}

// Go offers a redirection on "go <something>" but needs to be capitalized as a public symbol
func (c *TestMockExtensionCommands) Go(cmdArg string) string {
	return fmt.Sprintf("https://go.example.com/%s", url.QueryEscape(cmdArg))
}

// Fbgs offers a search of FullBigGrepSearch, as a convenience for the whole word (because 13 characters!)
func (c *TestMockExtensionCommands) Fbgs(cmdArg string) string {
	return c.FullBigGrepSearch(cmdArg)
}

// FullBigGrepSearch offers a search of the prescanned codebase "fbgs <something>" searches for <something>
func (c *TestMockExtensionCommands) FullBigGrepSearch(cmdArg string) string {
	return fmt.Sprintf("https://fullbiggrepsearch.example.com/pre-scanned/?q=%s", url.QueryEscape(cmdArg))
}
