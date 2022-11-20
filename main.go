package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

type ResponseMessage struct {
	Name   string
	Date   string
	Status int
}

func main() {
	repository_id := flag.String("id", "615bcf606feffc5384e8452e", "Container ID")
	flag.Parse()

	fmt.Println(TelegrafMessage(*repository_id))

}

func TelegrafMessage(id string) string {
	resp, err := http.Get("https://catalog.redhat.com/api/containers/v1/repositories/id/" + id)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	value := gjson.GetMany(string(body), "repository", "last_update_date")

	message := ResponseMessage{value[0].String(), value[1].String(), resp.StatusCode}
	bytemessage, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	return string(bytemessage)
}
