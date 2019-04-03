package skelton

import (
	"log"

	"github.com/urfave/cli"
)

func RunCli(args []string) {
	app := cli.NewApp()

	app.Name = "goskelton"
	app.Usage = "goskelton"

	app.Flags = flags
	app.Action = action

	app.Run(args)
}

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "project",
		Usage: "generated project name",
	},
	cli.StringFlag{
		Name:  "user",
		Usage: "git user name",
	},
	cli.StringFlag{
		Name:  "dest",
		Usage: "dest dir",
	},
}

func action(c *cli.Context) {
	config := &Config{
		Project: c.String("project"),
		User:    c.String("user"),
		Dest:    c.String("dest"),
	}

	if err := Run(config); err != nil {
		log.Fatal(err)
	}
}
