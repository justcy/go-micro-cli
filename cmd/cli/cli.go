// Package cli is a command line interface
package cli

import (
	"fmt"
	"os"
	osexec "os/exec"
	"strings"

	"github.com/chzyer/readline"
	"github.com/urfave/cli/v2"
	_ "github.com/justcy/go-micro/cmd/cli/new"
	_ "github.com/justcy/go-micro/cmd/cli/client"
)

var (
	prompt = "micro> "
)

func Run(c *cli.Context) error {
	// take the first arg as the binary
	binary := os.Args[0]

	r, err := readline.New(prompt)
	if err != nil {
		return err
	}
	defer r.Close()

	for {
		args, err := r.Readline()
		if err != nil {
			fmt.Fprint(os.Stdout, err)
			return err
		}

		args = strings.TrimSpace(args)

		// skip no args
		if len(args) == 0 {
			continue
		}

		parts := strings.Split(args, " ")
		if len(parts) == 0 {
			continue
		}

		cmd := osexec.Command(binary, parts...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(string(err.(*osexec.ExitError).Stderr))
		}
	}

	return nil
}

func init() {
}
