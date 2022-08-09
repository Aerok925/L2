package handler

import (
	"log"
	"net/http"
)

func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL, request.Method)
		if request.Method == http.MethodGet {
			log.Println()
		}
		next.ServeHTTP(writer, request)
	})
}
