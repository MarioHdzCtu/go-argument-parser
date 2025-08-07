package tests

import (
	"testing"

	namespace "github.com/MarioHdzCtu/argParser/src/namespace"
	app "github.com/MarioHdzCtu/argParser/src/parser"
)

func TestDefaultParser(t *testing.T) {
	ap := app.NewArgumentParser(app.ArgParserOptions{})

	if ap.Prog != "tests.test" {
		t.Errorf("Incorrect name in default parser: '%s'", ap.Prog)
	}

	if ap.PrefixChars != "-" {
		t.Errorf("Incorrect Prefix Chars '%s'", ap.PrefixChars)
	}

	if ap.AddHelp != true {
		t.Errorf("Incorrect value for Add Help '%t'", ap.AddHelp)
	}

	if ap.AllowAbbrev != true {
		t.Errorf("Incorrect value for Allow Abbrev '%t'", ap.AllowAbbrev)
	}

	if ap.ExitOnError != true {
		t.Errorf("Incorrect value for exit on error '%t'", ap.ExitOnError)
	}
}

func TestCustomParser(t *testing.T) {
	opts := app.ArgParserOptions{
		Prog:        "Test prog",
		Usage:       "Testing only",
		Description: "This is a test",
		Epilog:      "This is the end",
	}
	ap := app.NewArgumentParser(opts)

	if ap.Prog != opts.Prog {
		t.Errorf("Incorrect value for Prog. Expected '%s'. Got '%s'", opts.Prog, ap.Prog)
	}

	if ap.Usage != opts.Usage {
		t.Errorf("Incorrect value for Usage. Expected '%s'. Got '%s'", opts.Usage, ap.Usage)
	}

	if ap.Description != opts.Description {
		t.Errorf("Incorrect value for Description. Expected '%s'. Got '%s'", opts.Description, ap.Description)
	}

	if ap.Epilog != opts.Epilog {
		t.Errorf("Incorrect value for Epilog. Expected '%s'. Got '%s'", opts.Epilog, ap.Epilog)
	}
}

func TestPrintHelp(t *testing.T) {
	ap := app.NewArgumentParser(app.ArgParserOptions{Prog: "MytestProgram"})
	ap.PrintHelp()
}

func TestAutoHelp(t *testing.T) {
	ap := app.NewArgumentParser(app.ArgParserOptions{Prog: "My test Program"})

	if ap.Prog != "mytestprogram" {
		t.Errorf("Incorrect value for Prog")
	}

}

func TestNamespacePrint(t *testing.T) {
	ns := namespace.NewNamespace([]app.Argument{app.Argument{Name: "TestArg"}})
	ns.PrintNamespace()
}
