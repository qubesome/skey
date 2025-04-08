package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/qubesome/piv-go/piv"
	"github.com/urfave/cli/v3"
)

var (
	pivDevice string
)

func pivCommand() *cli.Command {
	cmd := &cli.Command{
		Name:  "piv",
		Usage: "Manage PIV Cards",
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List PIV Cards",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					cards, err := piv.Cards()
					if err != nil {
						return err
					}

					for _, card := range cards {
						fmt.Println(card)
					}
					return nil
				},
			},
			{
				Name:  "reset",
				Usage: "Factory reset PIV Cards",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:        "device",
						Min:         0,
						Max:         1,
						Destination: &pivDevice,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if pivDevice == "" {
						cards, err := piv.Cards()
						if err != nil {
							return err
						}

						if len(cards) == 0 {
							return fmt.Errorf("no PIV cards found")
						}

						if len(cards) > 1 {
							return fmt.Errorf("no device selected: multiple PIV cards found")
						}

						if len(cards) == 1 {
							pivDevice = cards[0]
						}
					}

					fmt.Printf("Selected PIV Card for FACTORY RESET: %q\n", pivDevice)
					if confirm("Are you sure you want to proceed? (Y to confirm):") {
						key, err := piv.Open(pivDevice)
						if err != nil {
							return fmt.Errorf("reset of %q failed: %w", pivDevice, err)
						}

						return key.Reset()
					}

					return nil
				},
			},
		},
	}
	return cmd
}

func confirm(msg string) bool {
	fmt.Print(msg + " ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return false
	}

	if strings.EqualFold(strings.TrimSpace(input), "Y") {
		return true
	} else {
		fmt.Println("Operation cancelled.")
	}
	return false
}
