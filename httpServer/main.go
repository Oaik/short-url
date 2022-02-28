package httpServer

import (
	"fmt"
	"net/http"
	"short-url/workerGenerator"
	"strings"
)

type Server struct {
	workerGen *workerGenerator.WorkerUrlGenerator
}

func New(workerGen *workerGenerator.WorkerUrlGenerator) *Server {
	w := Server{}
	w.workerGen = workerGen
	return &w
}

func (e *Server) StartServer() {
	http.HandleFunc("/short/", e.shortingFunc)
	http.HandleFunc("/goto/", e.redirectUrl)
	http.ListenAndServe(":7070", nil)
}

func (e *Server) shortingFunc(w http.ResponseWriter, req *http.Request) {
	websiteUrl := strings.TrimPrefix(req.URL.Path, "/short/")
	e.workerGen.AddWebsite(websiteUrl)
	fmt.Fprintf(w, websiteUrl)
}

func (e *Server) redirectUrl(w http.ResponseWriter, req *http.Request) {
	shortUrl := strings.TrimPrefix(req.URL.Path, "/goto/")
	websiteUrl := e.workerGen.SqlInstance.SelectWebsiteMatchShortUrl(shortUrl)
	if websiteUrl != "" {
		http.Redirect(w, req, "http://"+websiteUrl, 303)
		return
	}
	fmt.Fprintf(w, "Not Found")
}
