package cmdrepl

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

func ShellSplit(line string) (res []string, err error) {
	res = strings.Split(line, " ")
	err = nil
	return
}

// for{ prompt, readline, shellsplit, call app.Run() }
func CmdRepl(prompt string, app *cli.App) error {
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", prompt)
		line, ispr, err := rd.ReadLine()
		if err != nil {
			return err
		}
		lines := string(line)
		log.Println("readline", lines, ispr, err)
		tokens, err := ShellSplit(lines)
		if err != nil {
			return err
		}
		args := []string{tokens[0]}
		args = append(args, tokens...)
		log.Println("args", args)
		app.Run(args)
	}
	return nil
}
