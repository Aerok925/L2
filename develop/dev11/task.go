package main

import (
	"./cache"
	"./cache/cell"
	"./httpCalendar"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"net/http"
	"os"
)

/*
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1",  "date": "2017-01-01",  "event": "tester data 1"}' localhost:8080/create_event
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1",  "date": "2016-01-01",  "event": "tester data"}' localhost:8080/update_event
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1"}' localhost:8080/delete_event
*/

type conf struct {
	Addr string `yaml:"port"`
}

var (
	qwe httpCalendar.HTTPCalendar = httpCalendar.HTTPCalendar{
		Cache: cache.New(),
	}
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL, request.Method)
		qwe.Cache.Print()
		next.ServeHTTP(writer, request)
	})
}

func CreateJson(data interface{}) ([]byte, error) {
	retValue, err := json.Marshal(data)
	return retValue, err
}

func Create(data []byte) (*cell.Cell, int, error) {
	toCell, err := cell.ConvertToCell(data)
	if err != nil {
		return nil, 400, err
	}
	return toCell, 200, nil
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	all, _ := io.ReadAll(body)
	toCell, i, err := Create(all)
	if err != nil {
		data, _ := CreateJson("Error: " + err.Error())
		//w.Header().Set("Status code", strconv.Itoa(i))
		http.Error(w, err.Error(), i)
		fmt.Fprintln(w, string(data))
		return
	}
	err = qwe.Cache.LoadIn(toCell)
	if err != nil {
		data, _ := CreateJson("Error: " + err.Error())
		http.Error(w, err.Error(), 503)
		fmt.Fprintln(w, string(data))
		return
	}
	data, _ := CreateJson(toCell)
	fmt.Fprintln(w, string(data))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	all, _ := io.ReadAll(body)
	toCell, i, err := Create(all)
	if err != nil {
		data, _ := CreateJson("Error: " + err.Error())
		http.Error(w, err.Error(), i)
		fmt.Fprintln(w, string(data))
		return
	}
	err = qwe.Cache.Update(toCell)
	if err != nil {
		data, _ := CreateJson("Error: " + err.Error())
		http.Error(w, err.Error(), 503)
		fmt.Fprintln(w, string(data))
		return
	}
	data, _ := CreateJson(toCell)
	fmt.Fprintln(w, string(data))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	all, _ := io.ReadAll(body)
	toCell := cell.New()
	err := json.Unmarshal(all, toCell)
	if err != nil {
		data, _ := CreateJson("Error: " + err.Error())
		http.Error(w, err.Error(), 400)
		fmt.Fprintln(w, string(data))
		return
	}
	qwe.Cache.Delete(toCell.Uuid)
	data, _ := CreateJson(toCell)
	fmt.Fprintln(w, string(data))
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, qwe.Cache.Update([]byte("")))
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Incorrect number of arguments")
		return
	}
	confile := os.Args[1]
	file, err := os.Open(confile)
	if err != nil {
		log.Println(err)
		return
	}
	all, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	confS := &conf{}

	err = yaml.Unmarshal(all, confS)
	if err != nil {
		log.Println(err)
		return
	}

	getmux := http.NewServeMux()
	getmux.HandleFunc("/get_event", GetEvent)
	getmux.HandleFunc("/create_event", CreateEvent)
	getmux.HandleFunc("/update_event", UpdateEvent)
	getmux.HandleFunc("/delete_event", DeleteEvent)

	logHandler := MiddlewareLog(getmux)

	err = http.ListenAndServe(confS.Addr, logHandler)

}
