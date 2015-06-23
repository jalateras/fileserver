package main

import (
	"github.com/jalateras/fileserver/Godeps/_workspace/src/github.com/russross/blackfriday"
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8081", nil)
}

func GenerateMarkdown(resp http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	resp.Write(markdown)
}
