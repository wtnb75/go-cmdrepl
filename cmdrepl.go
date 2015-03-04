package cmdrepl

import (
	"bufio"
	"github.com/codegangsta/cli"
	"github.com/peterh/liner"
	"strings"
)

func ShellSplit(line string) (res []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	err = nil
	return
}

// for{ prompt, readline, shellsplit, call app.Run() }
func CmdRepl(prompt string, app *cli.App) error {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range app.Commands {
			if strings.HasPrefix(n.Name, line) {
				c = append(c, n.Name)
			}
		}
		return
	})
	for {
		if l, err := line.Prompt(prompt); err == nil {
			line.AppendHistory(l)
			tokens, err := ShellSplit(l)
			if err != nil {
				return err
			}
			args := []string{tokens[0]}
			args = append(args, tokens...)
			app.Run(args)
		} else {
			break
		}
	}
	return nil
}
