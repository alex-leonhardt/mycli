package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "greet",
		Usage:       "fight the loneliness!",
		Description: "....",
		Action: func(c *cli.Context) error {
			fmt.Println("hello friend!")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"MYCLI_CONFIG"},
				// Required: true,
				// FilePath: "/etc/mysql/password", // note, this takes precedence over env vars
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "task",
				Description: "task manages tasks",
				Subcommands: []*cli.Command{
					{
						Name:     "complete",
						Aliases:  []string{"c"},
						Category: "task",
						Usage:    "complete a task on the list",
						Action: func(c *cli.Context) error {
							fmt.Println("👏 you completed a task!")
							// return cli.Exit("nope", 1)
							return nil
							// return fmt.Errorf("nope")
						},
					},
					{
						Name:     "add",
						Aliases:  []string{"a"},
						Category: "task",
						Usage:    "add a task to the list",
						Action: func(c *cli.Context) error {
							fmt.Println("🚧 you added a task!")
							return nil
						},
					},
					{
						Name:     "edit",
						Aliases:  []string{"e"},
						Category: "task",
						Usage:    "edit a task on the list",
						Flags: []cli.Flag{
							// &cli.StringFlag{
							// 	Name:    "id",
							// 	Aliases: []string{"i"},
							// 	Usage:   "provide the task id to edit",
							// },
							&StdinFlag{
								Name:    "id",
								Aliases: []string{"i"},
							},
						},
						Action: func(c *cli.Context) error {
							fmt.Printf("🚧 %s is undergoing maintenance!\n", c.String("id"))
							return fmt.Errorf("")
							return nil
						},
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/*

























CUSTOM STDIN FLAG













*/

// StdinFlag is a custom flag that allows it to be set either by Stdin or by flag
type StdinFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool // when set, read from stdin will not work
	Hidden      bool
	TakesFile   bool
	Value       string
	DefaultText string
	Destination *string
	HasBeenSet  bool
}

// Apply populates the flag given the flag set and environment
func (f *StdinFlag) Apply(set *flag.FlagSet) error {

	var input []byte
	var err error
	var in os.FileInfo

	if in, _ = os.Stdin.Stat(); in.Mode()&os.ModeCharDevice == 0 {
		input, err = ioutil.ReadAll(os.Stdin)
		f.Value = strings.TrimSuffix(string(input), "\n")
		if err != nil {
			log.Fatal("booyaa")
		}
	}

	for _, name := range f.Names() {

		if f.Destination != nil {
			set.StringVar(f.Destination, name, f.Value, f.Usage)
			continue
		}
		set.String(name, f.Value, f.Usage)
	}

	return nil
}

// IsSet returns whether or not the flag has been set through env or file
func (f *StdinFlag) IsSet() bool {
	return f.HasBeenSet
}

// String returns a readable representation of this value
// (for usage defaults)
func (f *StdinFlag) String() string {
	return f.Value
}

// Names returns the names of the flag
func (f *StdinFlag) Names() []string {
	s := []string{f.Name}
	for _, v := range f.Aliases {
		s = append(s, v)
	}
	return s
}

// IsRequired returns whether or not the flag is required
func (f *StdinFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *StdinFlag) TakesValue() bool {
	return true
}

// GetUsage returns the usage string for the flag
func (f *StdinFlag) GetUsage() string {
	return f.Usage
}

// GetValue returns the flags value as string representation and an empty
// string if the flag takes no value at all.
func (f *StdinFlag) GetValue() string {
	return f.Value
}
