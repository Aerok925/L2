package cell

import (
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
