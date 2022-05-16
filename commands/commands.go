package commands

import (
	"fmt"
	"net/url"
	"regexp"
)

func init() {
	RegisterCommands(&Commands{})
}

var (
	m = regexp.MustCompile("^[A-Z]{3}-[0-9]{2,4}$")
)

// Commands is needed as a type at minimum to be able to pass to RegisterCommands()
type Commands struct{}

// Help and List are special-cased in Handler()
func (c *Commands) Help() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

// List and Help are special-cased in Handler()
func (c *Commands) List() {
	// We want this method to show up in `help`/`list` results
	// But the actual logic for these commands is in `main.go`
}

// G as a search prefix hands off to google
func (c *Commands) G(cmdArg string) string {
	return fmt.Sprintf("https://www.google.com/#q=%s", url.QueryEscape(cmdArg))
}

// Author retains a note to the original author
func (c *Commands) Author() string {
	return "https://www.markusdosch.com"
}

// So simply redirects the search to Stack Overflow's search URL
func (c *Commands) So(cmdArg string) string {
	return fmt.Sprintf("https://stackoverflow.com/search?q=%s", url.QueryEscape(cmdArg))
}

// TryRegex looks for a XXX-1234 -style pattern
func (c *Commands) TryRegex(parm string) string {
	if m.Match([]byte(parm)) {
		return fmt.Sprintf("https://two.example.com/%s", parm)
	}
	return ""
}
