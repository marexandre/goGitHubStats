package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"html/template"
	"log"
	"net/http"
)

var (
	userName = "marexandre"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(userName, opt)
	if err != nil {
		log.Panic(err)
	}

	var urls []string
	for _, v := range repos {
		urls = append(urls, github.Stringify(v.HTMLURL))
	}
	fmt.Printf("%v\n\n", urls)

	t, _ := template.ParseFiles("template/index.html", "template/default.html")
	err = t.ExecuteTemplate(w, "default", map[string]interface{}{"Urls": urls})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}
