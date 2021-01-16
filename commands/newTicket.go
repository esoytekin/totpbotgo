package commands

import (
	"github.com/esoytekin/totpbotgo/helpers"
	"github.com/esoytekin/totpbotgo/helpers/ajax"
	"github.com/esoytekin/totpbotgo/model"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

// NewTicketCommand command for creating new site
func NewTicketCommand() cli.Command {

	return cli.Command{
		Name:  "new",
		Usage: "creates new site data",
		Action: func(c *cli.Context) error {

			var secret string
			var ticket model.Ticket

			prompt := promptui.Prompt{Label: "enter site name"}

			siteName, err := prompt.Run()

			helpers.HandleError(err)

			prompt = promptui.Prompt{Label: "enter secret"}
			secret, err = prompt.Run()

			helpers.HandleError(err)

			ticket.Site = siteName
			ticket.Secret = secret

			createTicket(ticket)

			return nil

		},
	}

}

func createTicket(ticket model.Ticket) {

	ajax.Post(ticket)
}
