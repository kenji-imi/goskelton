package skelton

import (
	"log"

	"github.com/urfave/cli"
)

func RunCli(args []string) {
	app := cli.NewApp()

	app.Version = "0.0.1"
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
		Name:   "user",
		EnvVar: "GOSKELTON_USER",
		Usage:  "git user name for package import path in skelton file",
	},
	cli.StringFlag{
		Name:   "dest",
		Value:  ".",
		EnvVar: "GOSKELTON_DEST_DIR",
		Usage:  "path for under which directory to create project skelton",
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
