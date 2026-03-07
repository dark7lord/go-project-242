package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v3"
	"hexlet-path-size"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(context.Context, *cli.Command) error {
			result, err := path_size.GetPathSize("./testdata", false, false, false)
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
