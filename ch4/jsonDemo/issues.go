// Issues prints a table of Github issues matching the search terms.
package main

import (
	"fmt"
	"jsonDemo/githubissues"
	"log"
	"os"
)

func main() {
	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

/* Exercise 4.10: Modify issues to report the results in age categories, say less than a week old,
* less than an month old, less than a year old, and older than a year.
* Use the time package to compute the age of each issue. */

/* Exercise 4.11: Build a tool that lets users create, read, update, and delete GitHub issues
* from the command line, invoking their preferred text editor when substantial text input is required.
* Hint: Use the os/exec package to run the editor, and the encoding/json package to read */

/* Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
* https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
* many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
* that, using this index, prints the URL and transcript of each comic that matches a search term
* provided on the command line. */

/* Exercise 4.13: The JSON-based web service of the Open Movie Database lets you search
* https://omdbapi.com/ for a movie by name and download its poster image. Write a tool
* poster that downloads the poster image for the movie named on the command line. */
