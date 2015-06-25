package main

import (
	"net/http"
	"encoding/json"
	"html/template"
	"os"
	"fmt"
	"path"

	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/russross/blackfriday"
)

const (
	DEFAULT_PORT = "8080"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", HomeHandler)

	posts := router.Path("/posts").Subrouter()
	posts.Methods("GET").HandlerFunc(GetAllPostsHandler)
	posts.Methods("POSTS").HandlerFunc(CreateNewPostHandler)

	post := router.PathPrefix("/posts/{id}").Subrouter()
	post.Methods("GET").Path("/edit").HandlerFunc(EditPostHandler)
	post.Methods("GET").HandlerFunc(GetPostHandler)
	post.Methods("PUT").HandlerFunc(UpdatePostHandler)
	post.Methods("DELETE").HandlerFunc(DeletePostHandler)

	books := router.Path("/books").Subrouter()
	books.Methods("GET").HandlerFunc(GetAllBooksInHTML)
	server := negroni.Classic();
	server.UseHandler(router)
	server.Run(":" + port)
	fmt.Println("Server running on :" + port)
}

func GenerateMarkdown(resp http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	resp.Write(markdown)
}

func HomeHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Home")
}

func GetAllPostsHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Get All Posts")
}

func CreateNewPostHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Create New Post")
	resp.WriteHeader(201)
}

func EditPostHandler(resp http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Fprintln(resp, "Edit Post", id)
}

func GetPostHandler(resp http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Fprintln(resp, "Get Post", id)
}

func UpdatePostHandler(resp http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Fprintln(resp, "Update Post", id)
}

func DeletePostHandler(resp http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	fmt.Fprintln(resp, "Delete Post", id)
}

func GetAllBooks(resp http.ResponseWriter, req *http.Request) {
	book := Book {
		Title: "Building Web Apps with Go",
		Author: "Jeremy Saenz",
	}

	encoder := json.NewEncoder(resp)
	if err := encoder.Encode(book); err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
}


func GetAllBooksInHTML(resp http.ResponseWriter, req *http.Request) {
	book := Book {
		Title: "Building Web Apps with Go",
		Author: "Jeremy Saenz",
	}

	fp := path.Join("public", "templates", "books", "index.html.tmpl")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(resp, book); err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
}
