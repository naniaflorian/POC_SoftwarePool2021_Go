package main

import (
	"SofwareGoDay1/data"
	"SofwareGoDay1/humanity"
	"fmt"
	"log"
)

func main() {
	str, err := data.ReadFile("data/name.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := range str {
		oui, _ := data.LineToCSV(str[i])
		if ror != nil {
			log.Fatal(err)
		}
		humanity.NewHumanFromCSV(oui)
	}
}