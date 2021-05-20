// Example from book 'The go programming language'
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

var issueList = template.Must(template.New("issue_list").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HtmlURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HtmlURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HtmlUrl}}'>{{.Titel}}</a></td>
</tr>
{{end}}
</table>
`))

const IssueURL = "http://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HtmlURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string
}

type User struct {
	Login   string
	HtmlUrl string `json:"html_url"`
}

//
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	// text template report out
	/*
		report, err := template.New("report").
			Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// Text template
	report := template.Must(
		template.New("issueList").Funcs(template.FuncMap{"daysAgo": daysAgo}).
			Parse(templ))

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// Html template
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// Escape section
	var templText = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templText))

	var data struct {
		A string
		B template.HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
