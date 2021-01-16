package commands

import (
	"fmt"

	"github.com/esoytekin/totpbotgo/helpers"
	"github.com/esoytekin/totpbotgo/helpers/ajax"
	"github.com/urfave/cli"
)

// DeleteTicketCommand command to delete ticket
func DeleteTicketCommand() cli.Command {

	tickets := FetchTickets()

	return cli.Command{
		Name:  "delete",
		Usage: "delete site data",
		Action: func(c *cli.Context) error {

			ticket := helpers.ReadTicketID(c, tickets)
			if ticket == nil {
				return nil
			}

			ajax.Delete(ticket.ID)

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
