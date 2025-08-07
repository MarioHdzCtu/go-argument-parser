package namespace

import (
	"fmt"

	app "github.com/MarioHdzCtu/argParser/src/parser"
)

type Namespace struct {
	arguments []app.Argument
}

func (n *Namespace) PrintNamespace() {
	arguments_string := ""
	for _, arg := range n.arguments {
		arguments_string = arguments_string + arg.Name
	}
	fmt.Println(arguments_string)
}

func NewNamespace(args []app.Argument) *Namespace {
	ns := Namespace{arguments: args}

	return &ns
}
