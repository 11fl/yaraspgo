package yapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"yarasp/templates"
)

var zavidovo = "s9602593"

var moscow = "s2006004"

type Data struct {
	Segments []struct {
		Stops     string  `json:"stops,omitempty"`
		Departure string  `json:"departure,omitempty"`
		Arrival   string  `json:"arrival,omitempty"`
		Duration  float32 `json:"duration,omitempty"`
		Thread    struct {
			Number            string `json:"number,omitempty"`
			Transport_subtype struct {
				Title string `json:"title,omitempty"`
			} `json:"transport_subtype,omitempty"`
		} `json:"thread,omitempty"`
		From struct {
			Title string `json:"title,omitempty"`
		} `json:"from,omitempty"`
		To struct {
			Title string `json:"title,omitempty"`
		} `json:"to,omitempty"`
	} `json:"segments,omitempty"`
}

func ApiKey() string {
	key := os.Getenv("APIKEY")
	if key == "" {
		log.Fatalln("No api key provided")
	}
	return key
}

// url
func urlYa(f, t, d string) string {
	a := ApiKey()

	url := fmt.Sprintf("https://api.rasp.yandex.net/v3.0/search/?apikey=%s&format=json&from=%s&to=%s&lang=ru_RU&page=1&date=%s", a, f, t, d)
	return url

}

func fromto(from string) (t, f string) {
	if from == "zavidovo" {
		return zavidovo, moscow
	} else {
		return moscow, zavidovo
	}
}
func TimeTable(date, from string) string {
	f, t := fromto(from)
	url := urlYa(f, t, dates(date))
	d := Data{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
	}

	// fmt.Printf("response!\n")
	// fmt.Printf("status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("could not read response body: %s\n", err)
		os.Exit(1)
	}
	// fmt.Printf("client: response body: %s\n", resBody)

	err = json.Unmarshal(resBody, &d)
	if err != nil {
		log.Fatal(err)
	}

	funcMap := template.FuncMap{
		"Hours": Hours,
		//TODO: add function to format time to be more readable for human
		// "TFormat": TimeFormat,
	}
	var tempplaterasp bytes.Buffer

	tmp := template.New("simple").Funcs(funcMap)
	tmp, err = tmp.Parse(templates.RaspMessage)
	if err != nil {
		log.Fatal(err)
	}

	err = tmp.Execute(&tempplaterasp, &d)

	if err != nil {
		log.Fatal(err)
	}
	return tempplaterasp.String()
}

func Hours(seconds float32) string {
	s := int(seconds) % (24 * 3600)

	hour := s / 3600

	s %= 3600
	minutes := s / 60

	return fmt.Sprintf("%v:%v", hour, minutes)
}

func dates(date string) string {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	if date == "/today" {
		return fmt.Sprintf("%04d-%02d-%02d", today.Year(), today.Month(), today.Day())
	}
	if date == "/tomorrow" {
		return fmt.Sprintf("%04d-%02d-%02d", tomorrow.Year(), tomorrow.Month(), tomorrow.Day())
	}
	return ""
}
