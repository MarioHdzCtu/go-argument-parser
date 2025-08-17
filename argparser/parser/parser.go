package parser

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	argument "github.com/MarioHdzCtu/argParser/argparser/argument"
)

type ArgParserOptions struct {
	Prog                string
	Usage               string
	Description         string
	Epilog              string
	Parents             []argumentParser
	FormatterClass      string
	PrefixChars         string
	FromFilePrefixChars *os.File
	ArgumentDefault     any
	ConflictHandler     any //func
	AddHelp             *bool
	AllowAbbrev         *bool
	ExitOnError         *bool
}

type argumentParser struct {
	Prog                string
	Usage               string
	Description         string
	Epilog              string
	Parents             []argumentParser
	FormatterClass      string
	PrefixChars         string
	FromFilePrefixChars *os.File
	ArgumentDefault     any
	ConflictHandler     any //func
	AddHelp             bool
	AllowAbbrev         bool
	ExitOnError         bool
	arguments           []argument.Argument
	ParsedArguments     map[string]any
}

func (ap *argumentParser) PrintHelp() {
	fmt.Printf("Usage: %s [-h]\n", ap.Prog)
	fmt.Println("options:")
	fmt.Println("-h, --help\tshow this help message and exit")

	if len(ap.arguments) == 0 {
		fmt.Println(ap.Epilog)
		return
	}

	for _, arg := range ap.arguments {
		fmt.Printf("%s %s\t%s\n", arg.Name, arg.Flag, arg.Help)
	}

	fmt.Println()
	fmt.Println(ap.Epilog)

}

func (ap *argumentParser) AddArgument(a argument.Argument) {
	ap.arguments = append(ap.arguments, a)
}

func (ap *argumentParser) ParseArgs() *argumentParser {
	if len(os.Args) == 1 { // No arguments were passed to the script
		ap.PrintHelp()
		os.Exit(1)
	}

	if len(ap.arguments) >= len(os.Args)-1 { //Not enough arguments were passed to the script
		ap.PrintHelp()
		os.Exit(1)
	}

	for _, arg := range ap.arguments {

		idx_name := slices.Index(os.Args, arg.Name)
		idx_flag := 1

		if arg.Flag != "" {
			idx_flag = slices.Index(os.Args, arg.Flag)
		}

		if idx_name == -1 && idx_flag == -1 && arg.Required { //Argument is required but not present
			fmt.Printf("Argument %s is required\n", arg.Name)
			ap.PrintHelp()
			os.Exit(1)
		}

		correct_index := max(idx_name, idx_flag)
		value_to_cast := os.Args[correct_index+1]

		casted_value, err := dynamicCast(value_to_cast, arg.Vtype)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		striped_value := strings.ReplaceAll(arg.Name, "-", "")
		ap.ParsedArguments[striped_value] = casted_value
	}

	return ap
}

func dynamicCast(v string, targetType string) (any, error) {

	switch targetType {

	case "int":
		return strconv.Atoi(v)

	case "float":
		return strconv.ParseFloat(v, 64)

	case "string":
		return v, nil

	default:
		return nil, fmt.Errorf("unsupported target type: %s", targetType)

	}
}

func formatProg(prog string) (string, error) {
	noSpacesString := strings.ReplaceAll(prog, " ", "")
	lowerString := strings.ToLower(noSpacesString)

	return lowerString, nil
}

func NewArgumentParser(opts ArgParserOptions) *argumentParser {

	ap := argumentParser{
		Prog:            strings.Split(os.Args[0], "/")[len(os.Args)-1],
		PrefixChars:     "-",
		AddHelp:         true,
		AllowAbbrev:     true,
		ExitOnError:     true,
		ParsedArguments: make(map[string]any),
	}

	if opts.Prog != "" {
		ap.Prog, _ = formatProg(opts.Prog)
	}

	if opts.PrefixChars != "" {
		ap.PrefixChars = opts.PrefixChars
	}

	if opts.AddHelp != nil {
		ap.AddHelp = *opts.AddHelp
	}

	if opts.AllowAbbrev != nil {
		ap.AllowAbbrev = *opts.AllowAbbrev
	}

	if opts.ExitOnError != nil {
		ap.ExitOnError = *opts.ExitOnError
	}

	ap.Usage = opts.Usage
	ap.Description = opts.Description
	ap.Epilog = opts.Epilog
	ap.Parents = opts.Parents
	ap.FormatterClass = opts.FormatterClass
	ap.FromFilePrefixChars = opts.FromFilePrefixChars
	ap.ArgumentDefault = opts.ArgumentDefault
	ap.ConflictHandler = opts.ConflictHandler

	return &ap
}
