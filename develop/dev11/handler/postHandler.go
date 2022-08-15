package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Aerok925/L2/develop/dev11/cache"
	"github.com/Aerok925/L2/develop/dev11/cache/cell"
	"io/ioutil"
	"net/http"
	"time"
)

func parsPostMethod(r *http.Request) (*cell.Cell, error, int) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err, 503
	}
	defer r.Body.Close()
	tempCell := cell.New()
	err = json.Unmarshal(body, tempCell)
	if err != nil {
		return nil, err, 400
	}
	tempCell.DateTime, err = time.Parse("2006-01-2", tempCell.Date)
	if err != nil {
		return nil, err, 400
	}
	if tempCell.DateTime.Before(time.Now()) == true {
		return nil, errors.New("Old Date"), 400
	}
	if tempCell.Uuid == "" {
		return nil, errors.New(""), 400
	}
	return tempCell, nil, 0
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	tempCell, err, i := parsPostMethod(r)
	if err != nil {
		w.WriteHeader(i)
		fmt.Fprintln(w, string(NewErrRequest("Ошибка ввода данных")))
		return
	}
	err = cache.Storage.LoadIn(tempCell)
	if err != nil {
		w.WriteHeader(503)
		fmt.Fprintln(w, string(NewErrRequest(err.Error())))
	}
	fmt.Fprintln(w, string(NewReadyRequest("New event append!")))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	tempCell, err, i := parsPostMethod(r)
	if err != nil {
		w.WriteHeader(i)
		fmt.Fprintln(w, string(NewErrRequest("Ошибка ввода данных")))
		return
	}
	err = cache.Storage.Update(tempCell)
	if err != nil {
		w.WriteHeader(503)
		fmt.Fprintln(w, string(NewErrRequest(err.Error())))
		return
	}
	fmt.Fprintln(w, string(NewReadyRequest("Event update!")))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	tempCell, err, i := parsPostMethod(r)
	if err != nil {
		w.WriteHeader(i)
		fmt.Fprintln(w, string(NewErrRequest("Ошибка ввода данных")))
		return
	}
	err = cache.Storage.Delete(tempCell)
	if err != nil {
		w.WriteHeader(503)
		fmt.Fprintln(w, string(NewErrRequest("Не удалось удалить событие")))
		return
	}
	fmt.Fprintln(w, string(NewReadyRequest("Event deleted!")))
}
