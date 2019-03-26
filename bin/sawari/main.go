package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/sawari/bin/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sawari"
	app.Commands = []cli.Command{
		commands.Form,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
