package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/esoytekin/totpbotgo/helpers"
	"github.com/esoytekin/totpbotgo/model"

	"github.com/urfave/cli"
)

func TicketCommand(config model.Config) cli.Command { // {{{
	tickets := FetchTickets(config)
	return cli.Command{
		Name:  "ticket",
		Usage: "generates token for given ticket",
		Action: func(c *cli.Context) error {
			ticket := helpers.ReadTicketId(c, tickets)
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

} // }}}

var client = http.Client{Timeout: 10 * time.Second}

func FetchTickets(config model.Config) []model.Ticket {

	req, err := http.NewRequest("GET", config.Url, nil)

	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var result []model.Ticket

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func fetchToken(config model.Config, ticketId string) string {

	tokenUrl := fmt.Sprintf("%s/%s", config.Url, ticketId)

	req, err := http.NewRequest("GET", tokenUrl, nil)

	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	return string(result)
}
