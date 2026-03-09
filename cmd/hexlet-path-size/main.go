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
		Action: func(ctx context.Context, cmd *cli.Command) error {
			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")
			all := cmd.Bool("all")

			if cmd.NArg() == 0 {
				return fmt.Errorf("не указан путь")
			}
			filePath := cmd.Args().Get(0)

			size, err := path_size.GetPathSize(filePath, recursive, human, all)
			if err != nil {
				return fmt.Errorf("ошибка открытия по пути %s", filePath)
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
