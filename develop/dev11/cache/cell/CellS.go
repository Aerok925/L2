package cell

import (
	json2 "encoding/json"
	"time"
)

type Cell struct {
	Uuid     string
	Date     string
	Event    string
	DateTime time.Time `json:"-"`
}

func New() *Cell {
	return &Cell{}
}

func ConvertToCell(data []byte) (*Cell, error) {
	retCell := New()
	err := json2.Unmarshal(data, retCell)
	if err != nil {
		return nil, err
	}
	retCell.DateTime, err = time.Parse("2006-01-02", retCell.Date)
	if err != nil {
		return nil, err
	}
	return retCell, nil
}
