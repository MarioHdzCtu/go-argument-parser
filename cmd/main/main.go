package main

import (
	"fmt"

	app "github.com/MarioHdzCtu/argParser/src/parser"
)

func main() {
	ap := app.NewArgumentParser(app.ArgParserOptions{
		Prog:   "mycli",
		Epilog: "This is an epilog",
	})
	ap.AddArgument(app.Argument{Name: "--name", Vtype: "string", Required: true, Flag: "-n"})
	ap.AddArgument(app.Argument{Name: "--age", Vtype: "int", Required: true})

	ap = ap.ParseArgs()

	fmt.Println(ap.ParsedArguments["age"])
}
