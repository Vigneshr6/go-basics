package main

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"
)

var client *http.Client

func init() {
	log.Println("init client")
	client = http.DefaultClient
	client.Timeout = time.Second * 30
}

type Post struct {
	Id     int32  `json:"id"`
	UserId int32  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func TestHttpClient(t *testing.T) {
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("error getting json : %v", err)
	}
	log.Printf("http response : %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		panic("invalid response")
	}
	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		log.Fatalf("error parsing json : %v", err)
	}
	log.Printf("%+v", len(posts))
}
