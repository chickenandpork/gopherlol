package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TCommand is a mock empty set of commands with only an Author
type TCommand struct{}

// Author is as noted the only method in TCommands
func (c *TCommand) Author() string {
	return "test"
}

// TestLogger is a logsink that passes messages to t.Error
type TestLogger struct {
	t *testing.T
}

func (l TestLogger) Print(v ...interface{})                 { l.t.Log(v...) }
func (l TestLogger) Printf(format string, v ...interface{}) { l.t.Logf(format, v...) }
func (l TestLogger) Println(v ...interface{})               { l.t.Logf("%s", fmt.Sprintln(v...)) }

// TestRegister confirms that I didn't make an error in the simple singleton registry
func TestRegister(t *testing.T) {
	logsink = &TestLogger{t: t}

	initial := len(GetCommands()) // in case something else registers in func init()

	RegisterCommands(&TCommand{})
	RegisterCommands(&TCommand{})

	assert.Equal(t, 2, len(GetCommands())-initial)
}
