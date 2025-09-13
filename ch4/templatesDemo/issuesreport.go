// package templatesDemo
package main

import (
	"log"
	"os"
	"templatesDemo/githubissues"
	"text/template"
	"time"
)

// A template is a string or file containing one or more portions enclosed in double braces {{...}},
// called actions, that are evaluated at execution time to generate a text output.
const templ = `{{.TotalCount}} issues:
  {{range .Items}}
  Number: {{.Number}}
  User: {{.User.Login}}
  Title: {{.Title | printf "%.64s" }}
  Age: {{.CreatedAt | daysAgo }} days 
  {{end}}
`

// The {{range .Items}} and {{end}} actions create a loop, so the text between them
// is expanded multiple times, once for each element of the Items slice, with dot bound
// to successive elements of Items.
//
// Within an action, the | notation makes the result of one operation
// the argument of another, analogous to a Unix pipe.

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func reportGitHubIssues() {
	// func main() {
	// report, err := template.New("report").
	// 	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	// 	Parse(templ)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// The template.Must function is a helper that panics if the error is non-nil.
	// It simplifies safe initialization of global variables holding templates.
	report := template.Must(template.New("issueList").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))

	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
