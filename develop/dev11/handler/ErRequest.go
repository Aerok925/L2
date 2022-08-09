package handler

import "encoding/json"

type ErrRequest struct {
	Error string
}

func NewErrRequest(err string) []byte {
	temp := &ErrRequest{Error: err}
	marshal, _ := json.Marshal(temp)
	return marshal
}
