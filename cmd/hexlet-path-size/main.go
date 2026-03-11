// Package main implements a CLI utility to get the size
// of files and directories with output in a human-readable format.
package main

import (
	"context"
	"fmt"
	pathSize "hexlet-path-size"
	"log"
	"os"

	"github.com/urfave/cli/v3"
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
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")
			all := cmd.Bool("all")

			if cmd.NArg() == 0 {
				return fmt.Errorf("the path is not specified")
			}

			filePath := cmd.Args().Get(0)

			size, err := pathSize.GetPathSize(filePath, recursive, human, all)
			if err != nil {
				return fmt.Errorf("opening error along the %s path", filePath)
			}

			fmt.Println(size)

			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
