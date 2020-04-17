package commands

import (
	"fmt"
	"net/url"
)

func init() {
	RegisterCommands(&Commands{})
}

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
