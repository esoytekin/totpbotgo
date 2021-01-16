package helpers

import (
	"github.com/esoytekin/totpbotgo/model"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func getTicketByID(tickets []model.Ticket, ticketID string) *model.Ticket {
	for _, t := range tickets {
		if t.Site == ticketID {
			return &t
		}
	}

	return nil
}

// ReadTicketID returns ticket data for give ID
func ReadTicketID(c *cli.Context, tickets []model.Ticket) *model.Ticket {

	var ticketID string

	if c.NArg() > 0 {
		ticketID = c.Args().First()
		return getTicketByID(tickets, ticketID)
	}

	var ticketData []string

	for _, t := range tickets {
		ticketData = append(ticketData, t.Site)
	}

	prompt := promptui.Select{Label: "select ticket", Items: ticketData}

	index, _, err := prompt.Run()

	if err != nil {
		return nil
	}

	return &tickets[index]

}
