package humanity

import (
	"fmt"
)

type Human struct {
	Name string
	Age string
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	human := Human{Name: csv[0], Age: csv[1], Country: csv[2]}
	fmt.Printf("%+v\n", human)
	return nil, nil
}

//func NewHumansFromCsvFile(path string) ([]*Human, error)

