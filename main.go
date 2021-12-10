package main

import (
	"jrnl/cmd"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{

		Commands: []*cli.Command{cmd.TodayCmd},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
