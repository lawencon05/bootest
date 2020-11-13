package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "http://localhost:1234/"

type Request struct {
	Username string `json:"email"`
	Password string `json:"pwd"`
}

func main() {
	login()
}

func login() {
	jsonReq, err := json.Marshal(Request{
		Username: "test@gmail.com",
		Password: "123456",
	})
	url := baseURL + "api/user"

	var bearer = "Bearer xxxxxxx"

	// Create a new request using http
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response =>", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyBytes))
}
