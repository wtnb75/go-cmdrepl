package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/wtnb75/go-cmdrepl"
)

func hello(c *cli.Context) {
	fmt.Println("Hello World")
}

func ls(c *cli.Context) {
	fmt.Println("List World")
}

func main() {
	app := cli.NewApp()
	app.Name = "test"
	app.Usage = "Test"
	app.Author = ""
	app.Email = ""
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "boolflag,b",
			Usage: "testbool",
		},
		cli.IntFlag{
			Name:  "intvalue,i",
			Value: 10,
			Usage: "testint",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "ls",
			ShortName: "ls",
			Usage:     "list list list",
			Action:    ls,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "long,l",
					Usage: "long list",
				},
			},
		}, {
			Name:      "hello",
			ShortName: "hel",
			Usage:     "helooo",
			Action:    hello,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "morning,m",
					Usage: "good morning",
				},
			},
		},
	}
	cmdrepl.CmdRepl("test> ", app)
}
