package handler

import (
	"github.com/YKMeIz/logc"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	notFound, domain string
)

func init() {
	domain = os.Getenv("CARIBOU_DOMAIN")
	if domain == "" {
		logc.Error("environment variable CARIBOU_DOMAIN (e.g. google.ca) is not defined")
	}

	notFound = os.Getenv("CARIBOU_NOTFOUND")
	if domain == "" {
		logc.Error("environment variable CARIBOU_NOTFOUND (e.g. /404) is not defined")
	}

	domain = "//" + domain
	notFound = domain + notFound
}

func RootHandleFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// json api
		if strings.Contains(req.Header.Get("Accept"), "application/json") || strings.HasSuffix(req.URL.Path, ".json") {
			var norm bool
			u, e := url.Parse(req.Header.Get("Referer"))
			if e != nil {
				logc.Warning("error happened when parsing HTTP request header Referer: ", req.Header.Get("Referer"), "\n", e.Error())
			}
			if u.Host == domain[2:] {
				norm = true
			}

			if b := getMetadata(strings.TrimSuffix(filepath.Base(req.URL.Path), ".json"), norm); b != nil {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Content-Type", "application/json")
				w.Write(b)
				return
			}

			http.Redirect(w, req, notFound, http.StatusNotFound)
			return
		}

		// image short link
		if regexp.MustCompile(`\.[jpegnifJPEGNIF]{3,4}$`).MatchString(req.URL.Path) {
			url := getSourceURLByName(filepath.Base(req.URL.Path))
			if url == "" {
				http.Redirect(w, req, notFound, http.StatusNotFound)
				return
			}

			b, e := fetch(url)
			if e != nil {
				logc.Warning("error happened when pass https://i.pximg.net" + req.URL.Path + "\n" + e.Error())
				http.Redirect(w, req, notFound, http.StatusNotFound)
				return
			}

			w.Write(b)
			return
		}

		// serve file
		if _, err := os.Stat("dist" + req.URL.Path); os.IsNotExist(err) {
			w.Header().Set("Content-Encoding", "gzip")
			http.ServeFile(w, req, "dist/index.html")
			return
		}
		http.ServeFile(w, req, "dist"+req.URL.Path)

	}
}
