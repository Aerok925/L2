package handler

import "encoding/json"

type ReadyRequest struct {
	Result string
}

func NewReadyRequest(text string) []byte {
	temp := &ReadyRequest{Result: text}
	marshal, _ := json.Marshal(temp)
	return marshal
}
