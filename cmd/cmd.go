package cmd

import (
	"os"
	"path/filepath"
	"sort"
	"github.com/urfave/cli/v2"
)

var (
	version      = "0.0.1"
	description  = "A framework for cloud native development\n\n	 Use `mc [command] --help` to see command specific help."
	defaultFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "c",
			Usage:   "Set the config file: Defaults to ~/.micro/config.json",
			EnvVars: []string{"MICRO_CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Usage:   "Set the environment to operate in",
			EnvVars: []string{"MICRO_ENV"},
		},
	}
	app = cli.App{
		Name:      filepath.Base(os.Args[0]),
		HelpName:  filepath.Base(os.Args[0]),
		Usage:     description,
		Flags:     defaultFlags,
		Version:   version,
		Reader:    os.Stdin,
		Writer:    os.Stdout,
		ErrWriter: os.Stderr,
	}
)

func Register(cmds ...*cli.Command) {
	app.Commands = append(app.Commands, cmds...)
	sort.Slice(app.Commands, func(i, j int) bool {
		return app.Commands[i].Name < app.Commands[j].Name
	})
}

// Run the default command
func Run() {
	app.Run(os.Args)
}
