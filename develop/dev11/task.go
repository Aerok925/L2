package main

import (
	"./confg"
	"./handler"
	"log"
	"net/http"
	"os"
)

/*
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1",  "date": "2017-01-01",  "event": "tester data 1"}' localhost:8080/create_event
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1",  "date": "2016-01-01",  "event": "tester data"}' localhost:8080/update_event
curl -v -X POST -H "Content-Type: application/json" -d '{"uuid": "test1"}' localhost:8080/delete_event
*/

func main() {
	if len(os.Args) != 2 {
		log.Println("Incorrect number of arguments")
		return
	}
	// считываем конфигурационный файл
	conf := confg.New()
	err := conf.Load(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	getmux := http.NewServeMux()
	getmux.HandleFunc("/event_for_day", handler.EventsForDay)
	getmux.HandleFunc("/event_for_week", handler.EventsForWeek)
	getmux.HandleFunc("/event_for_month", handler.EventsForMonth)
	getmux.HandleFunc("/create_event", handler.CreateEvent)
	getmux.HandleFunc("/update_event", handler.UpdateEvent)
	getmux.HandleFunc("/delete_event", handler.DeleteEvent)

	logHandler := handler.MiddlewareLog(getmux)

	http.ListenAndServe(conf.Addr, logHandler)

}
