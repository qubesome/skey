package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/landlock-lsm/go-landlock/landlock"
	"github.com/qubesome/skey/cmd"
	"github.com/urfave/cli/v3"
)

var (
	debug bool
)

func main() {
	err := landlock.V5.BestEffort().Restrict(
		landlock.RODirs("/sys/devices"),   // qubesome/libudev: scanning FIDO devices
		landlock.RODirs("/run/udev/data"), // qubesome/libudev: scanning FIDO devices
		landlock.RODirs("/usr/lib64"),     // qubesome/piv-go: use of libpcsclite
	)
	if err != nil {
		fmt.Printf("failed to enforce landlock policies (requires Linux 5.13+): %w", err)
		os.Exit(1)
	}

	cmd := cmd.RootCommand()
	cmd.Flags = append(cmd.Flags, &cli.BoolFlag{
		Name:        "debug",
		Value:       false,
		Destination: &debug,
		Action: func(ctx context.Context, c *cli.Command, b bool) error {
			if debug {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
			return nil
		},
	})

	err = cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
