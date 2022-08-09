package handler

import (
	"../cache"
	"../cache/cell"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getDate(r *http.Request) (time.Time, error) {
	dateString := r.URL.Query().Get("date")
	dateTime, err := time.Parse("2006-01-2", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}

func getCells(r *http.Request, timer int) ([]cell.Cell, error) {
	uuid := r.URL.Query().Get("uuid")
	cells := cache.Storage.Get(uuid)
	dateTime, err := getDate(r)
	if err != nil {
		return nil, err
	}
	retCells := make([]cell.Cell, 0)
	tempTime := dateTime
	log.Println("date", dateTime.Year(), dateTime.Month(), dateTime.Day())
	for _, cell := range cells {
		dateTime = tempTime
		for i := 0; i < timer; i++ {
			if dateTime == cell.DateTime {
				retCells = append(retCells, cell)
			}
			dateTime = dateTime.Add(24 * time.Hour)
		}
	}
	//fmt.Print(tempTime.UTC())
	return retCells, nil
}

func createJSON(data interface{}) []byte {
	retValue := make([]byte, 0)
	retValue, _ = json.Marshal(data)

	return retValue
}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	retValue, err := getCells(r, 30)
	if err != nil {

		// сделать обработку ошибок
		return
	}
	data := createJSON(retValue)
	fmt.Fprintln(w, string(data))
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	retValue, err := getCells(r, 7)
	if err != nil {
		// сделать обработку ошибок
		return
	}
	data := createJSON(retValue)
	fmt.Fprintln(w, string(data))
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	retValue, err := getCells(r, 1)
	if err != nil {
		fmt.Fprintln(w, err)
		// сделать обработку ошибок
		return
	}
	data := createJSON(retValue)
	fmt.Fprintln(w, string(data))
}
