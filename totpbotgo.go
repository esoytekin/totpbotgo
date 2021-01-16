package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/esoytekin/totpbotgo/commands"
	"github.com/esoytekin/totpbotgo/model"

	"github.com/urfave/cli"
)

var config model.Config

func readConfiguration() {

	if config.Username != "" {
		return
	}

	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	finalPath := path.Join(usr.HomeDir, ".totpbotgo", "config.json")
	jsonFile, err := os.Open(finalPath)

	if err != nil {
		log.Fatal("couldn't find config file!")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

}

func createApp() *cli.App {
	app := cli.NewApp()
	app.Name = "totpbotgo"
	app.Usage = "generate otp token or qr code"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Emrah Soytekin",
			Email: "emrahsoytekin@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 Serious Enterprise"
	app.Version = "1.0.0"
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
	return app
}

func main() {

	app := createApp()

	readConfiguration()

	app.Commands = []cli.Command{
		commands.TicketCommand(config),
		commands.QrCommand(config),
		commands.NewTicketCommand(),
		commands.DeleteTicketCommand(),
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Dear %s, type help for usage information\n", config.Username)
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
