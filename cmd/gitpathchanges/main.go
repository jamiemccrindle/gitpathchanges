package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jamiemccrindle/gitpathchanges/pkg/gitpathchanges"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "The path to your git repository",
				Value:   ".",
			},
			&cli.StringSliceFlag{
				Name:    "directories",
				Aliases: []string{"d"},
				Usage:   "The directories to check for changes",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() != 2 {
				return fmt.Errorf("usage: gitpathchanges [options] commit1 commit2")
			}
			directories := c.StringSlice("directories")
			commitRef1 := c.Args().Get(0)
			commitRef2 := c.Args().Get(1)
			path := c.String("path")
			fmt.Println(gitpathchanges.Files(path, directories, commitRef1, commitRef2))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
