package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jamiemccrindle/gitpathchanges/pkg/gitpathchanges"
	"github.com/urfave/cli/v2"
)

const noResultsError = "no results"

func main() {
	app := &cli.App{
		Name:     "git-path-changes",
		HelpName: "git-path-changes",
		Usage:    "Check which files have changed between 2 commits",
		Authors: []*cli.Author{
			{
				Name:  "Jamie McCrindle",
				Email: "jamiemccrindle@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "match",
				Usage:     "List which matching files and directories have changed between 2 commits",
				ArgsUsage: "<commit1> <commit2>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "The path to your git repository",
						Value:   ".",
					},
					&cli.StringSliceFlag{
						Name:     "match",
						Aliases:  []string{"m"},
						Usage:    "The files or directories to check for changes",
						Required: true,
					},
				},

				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("usage: gitpathchanges [options] commit1 commit2")
					}
					matches := c.StringSlice("match")
					commitRef1 := c.Args().Get(0)
					commitRef2 := c.Args().Get(1)
					path := c.String("path")
					result, err := gitpathchanges.Files(path, matches, commitRef1, commitRef2)
					if err != nil {
						return err
					}
					for _, line := range *result {
						fmt.Println(line)
					}
					return nil
				},
			},
			{
				Name:      "has-matches",
				ArgsUsage: "<commit1> <commit2>",
				Usage:     "Return success if matching files have changed between 2 commits otherwise fail",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "The path to your git repository",
						Value:   ".",
					},
					&cli.StringSliceFlag{
						Name:     "match",
						Aliases:  []string{"m"},
						Usage:    "The files or directories to check for changes",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("usage: gitpathchanges [options] commit1 commit2")
					}
					matches := c.StringSlice("match")
					commitRef1 := c.Args().Get(0)
					commitRef2 := c.Args().Get(1)
					path := c.String("path")
					result, err := gitpathchanges.Files(path, matches, commitRef1, commitRef2)
					if err != nil {
						return err
					}
					if len(*result) == 0 {
						return fmt.Errorf(noResultsError)
					}
					return nil
				},
			},
			{
				Name:      "files",
				ArgsUsage: "<commit1> <commit2>",
				Usage:     "List all files that have changed between 2 commits",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "The path to your git repository",
						Value:   ".",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("usage: gitpathchanges [options] commit1 commit2")
					}
					commitRef1 := c.Args().Get(0)
					commitRef2 := c.Args().Get(1)
					path := c.String("path")
					result, err := gitpathchanges.Files(path, []string{}, commitRef1, commitRef2)
					if err != nil {
						return err
					}
					for _, line := range *result {
						fmt.Println(line)
					}
					return nil
				},
			},
			{
				Name:      "directories",
				ArgsUsage: "<commit1> <commit2>",
				Usage:     "List all files that have changed between 2 commits",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "path",
						Aliases: []string{"p"},
						Usage:   "The path to your git repository",
						Value:   ".",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() != 2 {
						return fmt.Errorf("usage: gitpathchanges [options] commit1 commit2")
					}
					commitRef1 := c.Args().Get(0)
					commitRef2 := c.Args().Get(1)
					path := c.String("path")
					result, err := gitpathchanges.Files(path, []string{}, commitRef1, commitRef2)
					if err != nil {
						return err
					}
					directories := make(map[string]struct{})
					for _, item := range *result {
						directories[gitpathchanges.Dirname(item)] = struct{}{}
					}
					for directory := range directories {
						fmt.Println(directory)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		if err.Error() == noResultsError {
			os.Exit(1)
			return
		}
		log.Fatal(err)
	}
}
