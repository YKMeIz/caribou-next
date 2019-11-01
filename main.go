package main

import (
	"github.com/YKMeIz/caribou/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user-profile/", handler.ProxyHandleFunc())
	mux.HandleFunc("/img-original/", handler.ProxyHandleFunc())
	mux.HandleFunc("/", handler.RootHandleFunc())

	http.ListenAndServe(":9090", mux)
}
