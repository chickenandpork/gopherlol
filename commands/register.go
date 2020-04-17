package commands

import (
	"errors"
	"sync"
)

// CommandSource defines the only required method: attribution of the commands source.  This is
// less about copyright or notoriety, more about "whom should I ask for details, help, support?"
type CommandSource interface {
	Author() string
}

// We singleton it here for utmost safety and code-simplicity in registering other command sources
// This leverages shameful copying of http://marcio.io/2015/07/singleton-pattern-in-go/
var (
	commandsInstance []CommandSource
	commandsOnce     sync.Once

	logsink LogSink = &DiscardingLogger{} // default to running quietly
)

// GetCommands is a singleton-enabling function: it returns the singleton, instantiating if necessary
func GetCommands() []CommandSource {
	if commandsInstance == nil {
		commandsOnce.Do(func() {
			commandsInstance = make([]CommandSource, 0)
			logsink.Println("allocating new COmmandSource[]")
		})
	}
	return commandsInstance
}

// GetCommands returns a value to side-step any write-backs
//func GetCommands() []CommandSource { s := getCommands() ; return *s; }

// RegisterCommands -- called directly or inside an `init()` function -- is used to register
// another source of commands by giving an instance of the class which defines those commands.  See
// `commands_test.go` for an example of the `init()` registration, or `register_test.go` for an
// example calling directly.
func RegisterCommands(c CommandSource) error {
	logsink.Println("registering new COmmandSource[]")
	if s := GetCommands(); s != nil {

		logsink.Println("COmmandSource[] len is ", len(commandsInstance))
		commandsInstance = append(commandsInstance, c)
		logsink.Print("COmmandSource[] len is ", len(commandsInstance))

		return nil
	}
	return errors.New("Commands Struct is Nil")
}
