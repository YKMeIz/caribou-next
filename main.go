package main

import (
	"github.com/YKMeIz/caribou/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/static/", handler.StaticHandleFunc())
	mux.HandleFunc("/user-profile/", handler.PassthroughHandleFunc())
	mux.HandleFunc("/img-original/", handler.PassthroughHandleFunc())
	mux.HandleFunc("/", handler.RootHandleFunc())

	http.ListenAndServe(":9090", mux)
}
