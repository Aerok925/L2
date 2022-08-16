package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
)

type wget struct {
	Addr     string
	download bool
}

func NewWget() *wget {
	retWget := &wget{}

	bo := flag.Bool("f", false, "in file")
	flag.Parse()
	if *bo {
		retWget.download = true
	}
	//retWget.download = *boolwget
	return retWget
}

func (w *wget) SetUrl(url string) error {
	_, err := url2.ParseRequestURI(url)
	if err != nil {
		return err
	}
	w.Addr = url
	return nil
}

func (w *wget) GetPage(c *http.Client) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, w.Addr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return all, nil
}

func (w *wget) Print(data []byte) {
	if w.download == false {
		log.Println(string(data))
		return
	}
	err := os.WriteFile("index.html", data, 0777)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	wg := NewWget()
	err := wg.SetUrl("https://linuxthebest.net/7-luchshih-emulyatorov-terminala-dlya-ubuntu-linux-mint/")
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	page, err := wg.GetPage(client)
	if err != nil {
		return
	}
	wg.Print(page)
}
