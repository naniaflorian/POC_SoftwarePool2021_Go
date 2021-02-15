package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(path string) ([]string, error) {

	content, err := ioutil.ReadFile(path)
	str := string(content)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(str, "\n"), err
}

func LineToCSV(line string) ([]string, error) {
	var err error
	str := strings.Split(line, ",")
	fmt.Printf("%s", str)
	return str, err
}

