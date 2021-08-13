package client

import (
	"fmt"
	"github.com/justcy/go-micro/cmd"
	"github.com/urfave/cli/v2"
)

var (
	defaultFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "-t",
			Usage:   "Set the template file: Defaults to ~/.micro/template",
			EnvVars: []string{"MICRO_NEW_TEMPLATE"},
		},
	}
)

type config struct {
	// foo
	Alias string
	// github.com/micro/foo
	Dir string
	// $GOPATH/src/github.com/micro/foo
	// go.micro.service.foo
	FQDN string

	GoDir string
	// $GOPATH
	GoPath string
	// UseGoPath
	UseGoPath bool
	// Files
	Files []file
	// Comments
	Comments []string
}

type file struct {
	Path string
	Tmpl string
}

func Run(ctx *cli.Context) error {
	fmt.Println("run mc client ...")
	return nil
}

func init() {
	cmd.Register(&cli.Command{
		Name:        "client",
		Usage:       "Create a client",
		Flags:       defaultFlags,
		Description: `'mc client' create a client for server . Example: 'mc client'`,
		Action:      Run,
	})
}
