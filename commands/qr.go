// Package commands provides ...
package commands

import (
	"fmt"
	"scratch/totpbotgo/helpers"
	"scratch/totpbotgo/model"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/urfave/cli"
)

func QrCommand(config model.Config) cli.Command {

	tickets := FetchTickets(config)

	return cli.Command{
		Name: "qrcode",
		Action: func(c *cli.Context) error {
			ticket := helpers.ReadTicketId(c, tickets)
			qrString := fmt.Sprintf("otpauth://totp/%s/:%s?secret=%s&issuer=%s", ticket.Site, config.Username, ticket.Secret, ticket.Site)
			obj := qrcodeTerminal.New()
			obj.Get(qrString).Print()
			return nil
		},
		BashComplete: func(c *cli.Context) {
			if c.NArg() > 0 {
				return
			}
			for _, t := range tickets {
				fmt.Println(t.Site)
			}
		},
	}
}
