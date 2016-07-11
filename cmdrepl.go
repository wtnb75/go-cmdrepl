package cmdrepl

import (
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/peterh/liner"
	"github.com/urfave/cli"
)

// CmdRepl : for{ prompt, readline, shellsplit, call app.Run() }
func CmdRepl(prompt string, app *cli.App) error {
	cli.OsExiter = func(rc int) {
		return
	}
	line := liner.NewLiner()
	defer line.Close()

	line.SetCompleter(func(line string) (c []string) {
		tokens, err := shellwords.Parse(line)
		if err != nil {
			return
		}
		lastnode := ""
		if len(tokens) != 0 {
			lastnode = tokens[len(tokens)-1]
		}
		nextprefix := strings.TrimSuffix(line, lastnode)
		// check command
		cmdlast := true
		cmdname := ""
		for i, n := range tokens {
			if !strings.HasPrefix(n, "-") {
				cmdname = n
				if i != len(tokens)-1 {
					cmdlast = false
				} else {
					cmdlast = true
				}
				break
			} else {
				cmdlast = false
			}
		}
		if cmdlast {
			// choose from Commands
			for _, n := range app.Commands {
				if strings.HasPrefix(n.Name, lastnode) {
					c = append(c, nextprefix+n.Name)
				}
				for _, nn := range n.Names() {
					if strings.HasPrefix(nn, lastnode) {
						c = append(c, nextprefix+nn)
					}
				}
			}
		} else {
			// choose from options
			flags := app.Flags
			for _, n := range app.Commands {
				if n.Name == cmdname {
					flags = n.Flags
					break
				}
				for _, nn := range n.Names() {
					if nn == cmdname {
						flags = n.Flags
						break
					}
				}
			}
			for _, n := range flags {
				for _, k := range strings.Split(n.GetName(), ",") {
					var flgstr string
					if len(k) == 1 {
						flgstr = "-" + k
					} else {
						flgstr = "--" + k
					}
					if strings.HasPrefix(flgstr, lastnode) {
						c = append(c, nextprefix+flgstr)
					}
				}
			}
		}
		return
	})
	for {
		if l, err := line.Prompt(prompt); err == nil {
			line.AppendHistory(l)
			tokens, err := shellwords.Parse(l)
			if err != nil {
				return err
			}
			if len(tokens) == 0 {
				continue
			}
			args := []string{tokens[0]}
			args = append(args, tokens...)
			if err = app.Run(args); err != nil {
				continue
			}
		} else {
			break
		}
	}
	return nil
}
