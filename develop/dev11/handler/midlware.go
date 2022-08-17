package handler

import (
	"log"
	"net/http"
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL, request.Method)
		next.ServeHTTP(writer, request)
	})
}
