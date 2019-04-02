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
		Name:  "name",
		Usage: "generated project name",
	},
	cli.StringFlag{
		Name:  "dest",
		Usage: "dest dir",
	},
}

func action(c *cli.Context) {
	config := &Config{
		Dest: c.String("dest"),
		Name: c.String("name"),
	}

	if err := Run(config); err != nil {
		log.Fatal(err)
	}
}
