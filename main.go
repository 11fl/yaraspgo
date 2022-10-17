package main

import (
	"log"
	"os"
	"time"
	"yarasp/telegram"
)

func ApiKey() string {
	key := os.Getenv("APIKEY")
	if key == "" {
		log.Fatalln("No api key provided")
	}
	return key
}

func main() {

	t := telegram.NewBot()
	for {
		t.GetUpdates()
		res := &t.Updates.Result
		if len(*res) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}

		t.UpdateResult(&t.Updates, &telegram.Offset{})
		time.Sleep(2 * time.Second)
	}
}
