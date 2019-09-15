package helpers

import (
	"scratch/totpbotgo/model"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func getTicketById(tickets []model.Ticket, ticketId string) *model.Ticket {
	for _, t := range tickets {
		if t.Site == ticketId {
			return &t
		}
	}

	return nil
}

func ReadTicketId(c *cli.Context, tickets []model.Ticket) *model.Ticket {
	var ticketId string
	if c.NArg() > 0 {
		ticketId = c.Args().First()
		return getTicketById(tickets, ticketId)
	} else {
		var ticketData []string

		for _, t := range tickets {
			ticketData = append(ticketData, t.Site)
		}

		prompt := promptui.Select{Label: "select ticket", Items: ticketData}

		index, _, err := prompt.Run()

		if err != nil {
			return nil
		} else {
			return &tickets[index]
		}
	}

}
