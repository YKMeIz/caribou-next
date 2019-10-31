package handler

import (
	"net/http"
	"os"
)

func StaticHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if _, err := os.Stat(req.URL.Path[1:]); os.IsNotExist(err) {
			http.Redirect(w, req, notFound, http.StatusNotFound)
			return
		}

		http.ServeFile(w, req, req.URL.Path[1:])
	}
}
