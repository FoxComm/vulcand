package command

import (
	"github.com/FoxComm/vulcand/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func NewLogCommand(cmd *Command) cli.Command {
	return cli.Command{
		Name: "log",
		Subcommands: []cli.Command{
			{
				ShortName: "set_severity",
				Usage:     "Set logging severity",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "severity, s"},
				},
				Action: cmd.updateLogSeverityAction,
			},
			{
				ShortName: "get_severity",
				Usage:     "Get logging severity",
				Action:    cmd.getLogSeverityAction,
			},
		},
	}
}

func (cmd *Command) updateLogSeverityAction(c *cli.Context) {
	sev := c.String("severity")
	if err := cmd.client.UpdateLogSeverity(sev); err != nil {
		cmd.printError(err)
		return
	}
	cmd.printOk("log severity updated")
}

func (cmd *Command) getLogSeverityAction(c *cli.Context) {
	sev, err := cmd.client.GetLogSeverity()
	if err != nil {
		cmd.printError(err)
		return
	}
	cmd.printOk("severity: %s", sev)
}
