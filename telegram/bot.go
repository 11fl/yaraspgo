package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"yarasp/fs"
	"yarasp/yapi"
)

var token = os.Getenv("tgtoken")
var apiUrl = fmt.Sprintf("https://api.telegram.org/bot%s", token)

type Bot struct {
	// Msg        Message
	PollingInt int
	Offset     int
	// Client     *http.Client
	Updates Updates
}

// TODO: ?
func NewBot() *Bot {
	newbot := &Bot{
		Offset: 0,
	}
	return newbot
}
func (b *Bot) SendMessage(chatid int, message string) {
	//TODO: move to other location
	url := apiUrl + "/sendMessage"

	msg := Message{
		Chat_Id: chatid,
		Text:    message,
	}

	json, err := json.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}

	body := bytes.NewBuffer([]byte(json))

	//TODO: make with timeout, move outside
	resp, err := http.Post(url, "application/json", body)

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Send message func :", resp.StatusCode, string(rbody))

}

func (b *Bot) GetUpdates() {
	var err error
	of := &Offset{Offset: b.Offset}
	currupdate := fs.ReadUpdate()

	of.Offset, err = strconv.Atoi(currupdate)

	if err != nil {
		log.Panic(err)
	}

	u := &b.Updates

	url := apiUrl + "/getUpdates"

	j, err := json.Marshal(of)
	if err != nil {
		log.Panic(err)
	}

	payload := bytes.NewBuffer(j)
	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		log.Panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(len(u.Result))

}

func (b *Bot) UpdateResult(u *Updates, of *Offset) {

	log.Println("Update result running")
	for r := range u.Result {
		of.Offset = u.Result[r].UpdateID + 1
		fs.WriteUpdate(fmt.Sprint(of.Offset))
		// fmt.Println(u.Result[r].Message.From.Id)
		// fmt.Println(u.Result[r].Message.From.Username)
		// fmt.Println(u.Result[r].Message.Text)
		fmt.Println(u.Result[r].UpdateID)
		fmt.Println(u.Ok)
		if u.Result[r].Message.Text == "/help" {

			id := u.Result[r].Message.From.Id
			text := "HELP"
			b.SendMessage(id, text)

		}
		if u.Result[r].Message.Text == "/start" {
			id := u.Result[r].Message.From.Id
			text := "Привет, человек!"
			b.SendMessage(id, text)
		}

		if u.Result[r].Message.Text == "/todayzavidovo" {
			tt := yapi.TimeTable("/todayzavidovo", "zavidovo")
			id := u.Result[r].Message.From.Id
			text := tt
			b.SendMessage(id, text)
		}
		if u.Result[r].Message.Text == "/tomorrowzavidovo" {
			tt := yapi.TimeTable("/tomorrowzavidovo", "zavidovo")
			id := u.Result[r].Message.From.Id
			text := tt
			b.SendMessage(id, text)
		}
		if u.Result[r].Message.Text == "/todaymoscow" {
			tt := yapi.TimeTable("/todaymoscow", "moscow")
			id := u.Result[r].Message.From.Id
			text := tt
			b.SendMessage(id, text)
		}
		if u.Result[r].Message.Text == "/tomorrowmoscow" {
			tt := yapi.TimeTable("/tomorrowmoscow", "moscow")
			id := u.Result[r].Message.From.Id
			text := tt
			b.SendMessage(id, text)
		}
	}
}
