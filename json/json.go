package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var jsonData []byte

func init() {
	data, err := ioutil.ReadFile("./article.json")
	if err != nil {
		panic("error reading file")
	}
	jsonData = data
}

type articleData struct {
	Data []struct {
		DataObj
		Attributes    `json:"attributes"`
		Relationships struct {
			Author struct {
				Data DataObj `json:"data"`
			} `json:"author"`
		} `json:"relationships"`
	} `json:"data"`
}

type DataObj struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Attributes struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Created MyTime `json:"created"`
	Updated MyTime `json:"updated"`
}

type MyTime struct {
	time.Time
}

func (t *MyTime) UnmarshalJSON(b []byte) error {
	timeString := strings.TrimSpace(string(b))
	if timeString == "" {
		return nil
	}
	// tt, err := time.Parse(time.RFC3339Nano, timeString)
	// 2015-05-22T14:56:29.000Z
	tt, err := time.Parse("\"2006-01-02T15:04:05.999Z\"", timeString)
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

func (t *MyTime) MarshalJSON() ([]byte, error) {
	//TODO
	return nil, nil
}

func main() {
	var a articleData
	if err := json.Unmarshal(jsonData, &a); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", a)
}
