package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestGetRequest(t *testing.T) {
	urlAddr := "http://127.0.0.1:8080/create_event"
	transport := &http.Transport{}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	testerFiles := []string{"json_tests/create_event/PostValidData.json"}
	open, err := os.Open(testerFiles[0])
	if err != nil {
		log.Println(err)
		return
	}
	defer open.Close()
	data, err := ioutil.ReadAll(open)
	if err != nil {
		return
	}
	body := bytes.NewBufferString(string(data))
	req, err := http.NewRequest(http.MethodPost, urlAddr, body)
	if err != nil {
		log.Println(err)
		return
	}
	req.
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	log.Println(string(respBody))
	log.Println(resp.StatusCode)
}
