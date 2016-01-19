package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

var issueList = template.Must(template.New("issueList").Parse(`
<h1>{{ .TotalCount }}</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{ range .Items }}
<tr>
	<td><a href='{{ .HTMLURL }}'>{{ .Number }}</a></td>
	<td>{{ .State }}</td>
	<td><a href='{{ .User.HTMLURL }}'>{{ .User.Login }}</a></td>
	<td><a href='{{ .HTMLURL }}'>{{ .Title }}</a></td>
</tr>
{{ end }}
</table>
`))

func Handler(w http.ResponseWriter, r *http.Request) {
	params := []string{"repo:golang/go", "3133", "10535"}
	result, err := SearchIssues(params)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	if err := issueList.Execute(w, result); err != nil {
		fmt.Fprintf(w, "%v", err)
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
