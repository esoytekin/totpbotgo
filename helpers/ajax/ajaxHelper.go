package ajax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/esoytekin/totpbotgo/model"
)

var config model.Config

var client = http.Client{Timeout: 10 * time.Second}

// Fetch send get request
func Fetch() []model.Ticket {

	readConfiguration()

	var result []model.Ticket

	req, err := http.NewRequest("GET", config.URL, nil)

	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Post send post request
func Post(ticket model.Ticket) {
	readConfiguration()

	ticketBody, err := json.Marshal(&ticket)
	req, err := http.NewRequest("POST", config.URL, bytes.NewReader(ticketBody))

	req.SetBasicAuth(config.Username, config.Password)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

}

// Delete sends delete request
func Delete(ticketID int64) {
	readConfiguration()

	endURL := fmt.Sprintf("%s/%d", config.URL, ticketID)

	req, err := http.NewRequest("DELETE", endURL, nil)

	req.SetBasicAuth(config.Username, config.Password)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

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
