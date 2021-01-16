package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/esoytekin/totpbotgo/helpers"
	"github.com/esoytekin/totpbotgo/helpers/ajax"
	"github.com/esoytekin/totpbotgo/model"

	"github.com/urfave/cli"
)

// TicketCommand command for printing the token for given site
func TicketCommand(config model.Config) cli.Command {
	tickets := FetchTickets()
	return cli.Command{
		Name:  "ticket",
		Usage: "generates token for given ticket",
		Action: func(c *cli.Context) error {
			ticket := helpers.ReadTicketID(c, tickets)
			if ticket == nil {
				return nil
			}
			token := fetchToken(config, ticket.Site)
			fmt.Println(token)
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

var client = http.Client{Timeout: 10 * time.Second}

// FetchTickets returns all available ticket list
func FetchTickets() []model.Ticket {

	return ajax.Fetch()
}

func fetchToken(config model.Config, ticketID string) string {

	tokenURL := fmt.Sprintf("%s/%s", config.URL, ticketID)

	req, err := http.NewRequest("GET", tokenURL, nil)

	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	return string(result)
}
