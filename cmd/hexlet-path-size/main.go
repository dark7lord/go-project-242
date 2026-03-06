package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"os"
	"log"
	"fmt"
	"hexlet-path-size/sizegetters"
)

func main() {
	cmd := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(context.Context, *cli.Command) error {
			result, err := sizegetters.GetPathSize("./mocks/tmnt.csv", false, false, false)

			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}
	
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal()
	}
}
