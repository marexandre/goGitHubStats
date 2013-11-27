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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(userName, opt)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n\n", github.Stringify(repos))

	projects := []Project{}
	for _, p := range repos {
		project := Project{}
		project.ID = p.ID
		project.URL = p.URL
		project.HTMLURL = p.HTMLURL
		project.Name = p.Name
		project.Description = p.Description
		projects = append(projects, project)
	}

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
