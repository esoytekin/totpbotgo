// Package commands provides ...
package commands

import (
	"fmt"

	"github.com/esoytekin/totpbotgo/helpers"
	"github.com/esoytekin/totpbotgo/model"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/urfave/cli"
)

// QrCommand prints qr code for givin site
func QrCommand(config model.Config) cli.Command {

	tickets := FetchTickets()

	return cli.Command{
		Name:  "qrcode",
		Usage: "generates qr code for given ticket",
		Action: func(c *cli.Context) error {
			ticket := helpers.ReadTicketID(c, tickets)
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
