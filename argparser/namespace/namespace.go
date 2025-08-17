package namespace

import (
	"fmt"

	argument "github.com/MarioHdzCtu/argParser/argparser/argument"
)

type Namespace struct {
	arguments []argument.Argument
}

func (n *Namespace) PrintNamespace() {
	arguments_string := ""
	for _, arg := range n.arguments {
		arguments_string = arguments_string + arg.Name
	}
	fmt.Println(arguments_string)
}

func NewNamespace(args []argument.Argument) *Namespace {
	ns := Namespace{arguments: args}

	return &ns
}
