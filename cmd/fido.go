package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/qubesome/libudev"
	"github.com/qubesome/libudev/matcher"
	"github.com/urfave/cli/v3"
)

func fidoCommand() *cli.Command {
	cmd := &cli.Command{
		Name: "fido",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List FIDO keys",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					m := matcher.NewMatcher()
					m.SetStrategy(matcher.StrategyOr)
					m.AddRule(matcher.NewRuleEnv("ID_FIDO_TOKEN", "1"))
					m.AddRule(matcher.NewRuleEnv("ID_SECURITY_TOKEN", "1"))

					scanner, err := libudev.NewScanner(
						libudev.WithMatcher(m),
						libudev.WithPathFilterPattern(regexp.MustCompile("(?i)^.*pci0000:00.*usb.*")),
					)

					if err != nil {
						return err
					}

					devices, err := scanner.ScanDevices()
					if err != nil {
						return err
					}

					for _, d := range devices {
						vendor := "N/A"
						product := "N/A"
						if d.Parent != nil {
							vendor = "0x" + d.Parent.VendorID
							product = "0x" + d.Parent.ProductID
						}

						slog.Debug(d.Env["HID_NAME"], "devpath", d.Devpath, "tags", d.Tags)
						fmt.Printf("/dev/%s: vendor=%s, product=%s (%s)\n",
							d.Env["DEVNAME"],
							vendor,
							product,
							d.Parent.Env["HID_NAME"],
						)
					}

					return nil
				},
			},
		},
	}
	return cmd
}
