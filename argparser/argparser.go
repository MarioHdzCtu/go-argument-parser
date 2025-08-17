package argparser

import (
	argument "github.com/MarioHdzCtu/argParser/argparser/argument"
	app "github.com/MarioHdzCtu/argParser/argparser/parser"
)

var NewArgumentParser = app.NewArgumentParser

var NewArgument = argument.NewArgument

func BoolP(b bool) *bool {
	return &b
}

type ArgParserOptions = app.ArgParserOptions

type ArgumentOptions = argument.ArgumentOptions
