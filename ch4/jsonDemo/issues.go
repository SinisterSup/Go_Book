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
