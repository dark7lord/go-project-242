package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"os"
	"log"
	"fmt"
)

func main() {
	cmd := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}
	
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal()
	}
}
