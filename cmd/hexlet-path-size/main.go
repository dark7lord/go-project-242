package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"hexlet-path-size"
	"log"
	"os"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			human := cmd.Bool("human")

			result, err := path_size.GetPathSize("./testdata", false, human, false)
			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal()
	}
}
