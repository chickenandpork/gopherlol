package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TCommand is a mock empty set of commands with only an Author
type TCommand struct{}

// Author is as noted the only method in TCommands
func (c *TCommand) Author() string {
	return "test"
}

// TryRegex is a method of checking a regex
func (c *TCommand) TryRegex(parm string) string {
	return ""
}

// TestRegister confirms that I didn't make an error in the simple singleton registry
func TestRegister(t *testing.T) {
	assert.Nil(t, SetLogSink(t))

	initial := len(GetCommands()) // in case something else registers in func init()

	RegisterCommands(&TCommand{})
	RegisterCommands(&TCommand{})

	assert.Equal(t, 2, len(GetCommands())-initial)
}
