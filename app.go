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

type Project struct {
	ID          *int
	URL         *string
	HTMLURL     *string
	Name        *string
	Description *string
}

func filterProjects(rs []github.Repository) []Project {
	ps := []Project{}
	for _, v := range rs {
		p := Project{}
		p.ID = v.ID
		p.URL = v.URL
		p.HTMLURL = v.HTMLURL
		p.Name = v.Name
		p.Description = v.Description
		ps = append(ps, p)
	}
	return ps
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(userName, opt)

	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n\n", github.Stringify(repos))

	projects := filterProjects(repos)

	t, _ := template.ParseFiles("template/base.html", "template/index.html")
	err = t.ExecuteTemplate(w, "base", map[string]interface{}{"Projects": projects})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}
