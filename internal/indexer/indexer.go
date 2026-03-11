package indexer


import "fmt"

func Index(page string, links []string) {

	fmt.Println("Indexed:", page)
	fmt.Println("Links found:", len(links))

}