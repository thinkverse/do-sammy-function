package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ENDPOINT = "https://functionschallenge.digitalocean.com/api/sammy"

type Sammy struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Response struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Body       string `json:"body,omitempty"`
}

func Main(args map[string]interface{}) *Response {
	b, _ := json.Marshal(Sammy{
		Name: "thinkverse",
		Type: "punk",
	})

	req, err := http.NewRequest(http.MethodPost, ENDPOINT, bytes.NewBuffer(b))

	if err != nil {
		panic(err)
	}

	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return &Response{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}
}
