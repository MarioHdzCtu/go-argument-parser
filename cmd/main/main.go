package main

import (
	"fmt"

	argparser "github.com/MarioHdzCtu/argParser/argparser"
)

func main() {
	ap := argparser.NewArgumentParser(argparser.ArgParserOptions{
		Prog:   "mycli",
		Epilog: "This is an epilog",
	})

	ap.AddArgument(*argparser.NewArgument(argparser.ArgumentOptions{Name: "--name", Vtype: "string", Required: argparser.BoolP(true)}))
	ap.AddArgument(*argparser.NewArgument(argparser.ArgumentOptions{Name: "--age", Vtype: "int", Required: argparser.BoolP(true)}))
	ap = ap.ParseArgs()

	fmt.Println(ap.ParsedArguments["age"])
}
