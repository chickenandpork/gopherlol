package commands

import (
	"errors"
	"fmt"

	//"net/url"
	"sync"
)

type CommandSource interface {
	Author() string
}

// We singleton it here for utmost safety and code-simplicity in registering other command sources
// This leverages shameful copying of http://marcio.io/2015/07/singleton-pattern-in-go/
var (
	commandsInstance []CommandSource
	commandsOnce     sync.Once
)

// getCommands is a singleton-enabling function: it returns the singleton, instantiating if necessary
func GetCommands() []CommandSource {
	if commandsInstance == nil {
		commandsOnce.Do(func() {
			commandsInstance = make([]CommandSource, 0)
			fmt.Println("allocating new COmmandSource[]")
		})
	}
	return commandsInstance
}

// GetCommands returns a value to side-step any write-backs
//func GetCommands() []CommandSource { s := getCommands() ; return *s; }

func RegisterCommands(c CommandSource) error {
	fmt.Println("registering new COmmandSource[]")
	if s := GetCommands(); s != nil {

		fmt.Println("COmmandSource[] len is ", len(commandsInstance))
		commandsInstance = append(commandsInstance, c)
		fmt.Print("COmmandSource[] len is ", len(commandsInstance))

		return nil
	} else {
		return errors.New("Commands Struct is Nil")
	}
}
