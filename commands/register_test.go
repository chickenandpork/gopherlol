package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TCommand struct{}

func (c *TCommand) Author() string {
	return "test"
}

// TestRegister confirms that I didn't make an error in the simple singleton registry
func TestRegister(t *testing.T) {

	initial := len(GetCommands()) // in case something else registers in func init()

	RegisterCommands(&TCommand{})
	RegisterCommands(&TCommand{})

	assert.Equal(t, 2, len(GetCommands())-initial)
}
