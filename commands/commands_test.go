package commands

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHandler codifies the expected behavior of the Handler() to ensure that the freedom to
// change/refactor/extend does not break current functionality
func TestHandler(t *testing.T) {
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
				"<li><strong>so</strong>, takes args</li></ul>"},
		{
			"list",
			"/?q=help", http.StatusOK,
			"<h1>gopherlol command list</h1><ul>" +
				"<li><strong>author</strong></li>" +
				"<li><strong>g</strong>, takes args</li>" +
				"<li><strong>help</strong></li>" +
				"<li><strong>list</strong></li>" +
				"<li><strong>so</strong>, takes args</li></ul>"},
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
			assert.Equalf(t, expected, observed, "handler returned incorrect text: expected %v observed %v", expected, observed)
		})
	}
}
