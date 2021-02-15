package main

import (
	"SofwareGoDay1/data"
	"fmt"
	"log"
)

func main() {
	str, err := data.ReadFile("data/name.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
	for i, v := range str {
		fmt.Printf("%d", i)
		data.LineToCSV(v)
	}
}