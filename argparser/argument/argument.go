package argument

import (
	"fmt"
	"os"

	utils "github.com/MarioHdzCtu/argParser/argparser/utils"
)

type Argument struct {
	Name       string
	Flag       string
	Nargs      int
	Vconst     string
	Vdefault   any
	Vtype      string
	Choices    []any
	Required   bool
	Help       string
	Metavar    string
	Dest       string
	Deprecated bool
}

type ArgumentOptions struct {
	Name       string
	Flag       string
	Nargs      *int
	Vconst     any
	Vdefault   any
	Vtype      string
	Choices    []any
	Required   *bool
	Help       string
	Metavar    string
	Dest       string
	Deprecated bool
}

func NewArgument(opts ArgumentOptions) *Argument {

	argument := Argument{
		Nargs: 1,
		Vtype: "string",
	}

	//logic checks
	if opts.Name == "" && opts.Flag == "" {
		fmt.Println("Either the name or flag should be provided to the Argument")
		os.Exit(1)
	}

	if *opts.Required == false && opts.Vdefault != nil {
		fmt.Println("If field is not required a default value should be provided")
	}

	//default assignments

	if opts.Name != "" {
		argument.Name = opts.Name
	}

	if opts.Flag != "" {
		argument.Flag = opts.Flag
	}

	if opts.Nargs != nil {
		argument.Nargs = *opts.Nargs
	}

	//add Vconst

	if opts.Vdefault != nil {
		argument.Vdefault = opts.Vdefault
	}

	if opts.Vtype != "" {
		argument.Vtype = opts.Vtype
	}

	if opts.Choices != nil {
		argument.Choices = opts.Choices
	}

	if opts.Required != nil {
		argument.Required = *opts.Required
	}

	if opts.Help != "" {
		argument.Help = opts.Help
	}

	if opts.Metavar == "" {
		v, err := utils.Coalesce([]any{opts.Name, opts.Flag})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if s, ok := v.(string); ok {
			argument.Metavar = s
		} else {
			fmt.Println("Error on returned value type for metavar")
			os.Exit(1)
		}
	}

	return &argument
}
