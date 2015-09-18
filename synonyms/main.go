package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/onufert/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thes := &thesaurus.BigHugeLabs{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thes.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms")
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonmys")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
