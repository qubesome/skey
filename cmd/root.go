package cmd

import (
	"github.com/urfave/cli/v3"
)

func RootCommand() *cli.Command {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			fidoCommand(),
			pivCommand(),
		},
	}

	cmd.Usage = "A Go Native cli to manage secure keys"
	cmd.Suggest = true
	cmd.EnableShellCompletion = true

	return cmd
}
