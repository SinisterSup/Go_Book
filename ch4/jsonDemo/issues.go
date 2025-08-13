// Issues prints a table of Github issues matching the search terms.
package main

import (
	"fmt"
	"jsonDemo/githubissues"
	"log"
	"os"
	"time"
)

//
// func main() {
// 	result, err := githubissues.SearchIssues(os.Args[1:])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Printf("%d issues:\n", result.TotalCount)
// 	for _, item := range result.Items {
// 		fmt.Printf("#%-5d %9.9s %.55s\n",
// 			item.Number, item.User.Login, item.Title)
// 	}
// }

/* Exercise 4.10: Modify issues to report the results in age categories, say less than a week old,
* less than an month old, less than a year old, and older than a year.
* Use the time package to compute the age of each issue. */
func modifiedIssues() {
	// func main() {
	result, err := githubissues.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %d issues\n", result.TotalCount)

	// Define the age categories
	var lessThanWeek, lessThanMonth, lessThanYear, moreThanYear []*githubissues.Issue

	// Current time for age comparison
	now := time.Now()

	// Categorize issues by age
	for _, item := range result.Items {
		// Parse the creation time
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			log.Printf("Error parsing time for issue #%d: %v", item.Number, err)
			continue
		}

		// Calculate age and categorize
		age := now.Sub(createdAt)

		switch {
		case age < 7*24*time.Hour:
			lessThanWeek = append(lessThanWeek, item)
		case age < 30*24*time.Hour:
			lessThanMonth = append(lessThanMonth, item)
		case age < 365*24*time.Hour:
			lessThanYear = append(lessThanYear, item)
		default:
			moreThanYear = append(moreThanYear, item)
		}
	}

	// Print issues by category with counts
	fmt.Printf("\n===== Issues less than a week old (%d) =====\n", len(lessThanWeek))
	printIssues(lessThanWeek)

	fmt.Printf("\n===== Issues less than a month old (%d) =====\n", len(lessThanMonth))
	printIssues(lessThanMonth)

	fmt.Printf("\n===== Issues less than a year old (%d) =====\n", len(lessThanYear))
	printIssues(lessThanYear)

	fmt.Printf("\n===== Issues more than a year old (%d) =====\n", len(moreThanYear))
	printIssues(moreThanYear)
}

// printIssues prints a formatted list of issues
func printIssues(issues []*githubissues.Issue) {
	if len(issues) == 0 {
		fmt.Println("No issues in this category")
		return
	}

	for _, item := range issues {
		// Parse time to display the actual age
		createdAt, _ := time.Parse(time.RFC3339, item.CreatedAt)
		age := time.Since(createdAt)

		// Format age in a human-readable format
		var ageStr string
		switch {
		case age < 24*time.Hour:
			ageStr = fmt.Sprintf("%.0f hours", age.Hours())
		case age < 30*24*time.Hour:
			ageStr = fmt.Sprintf("%.0f days", age.Hours()/24)
		case age < 365*24*time.Hour:
			ageStr = fmt.Sprintf("%.1f months", age.Hours()/(24*30))
		default:
			ageStr = fmt.Sprintf("%.1f years", age.Hours()/(24*365))
		}

		fmt.Printf("#%-5d %9.9s %8s %.55s\n",
			item.Number, item.User.Login, ageStr, item.Title)
	}
}

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
