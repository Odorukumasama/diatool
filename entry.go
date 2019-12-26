package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type entry struct {
	id   int
	date time.Time
}

func NewEntry(id int) *entry {
	e := new(entry)
	e.id = id

	e.date = time.Now()

	return e
}

func main() {
	eintrag1 := NewEntry(1)
	fmt.Printf("Id eintrag1 lautet: %v\n", eintrag1.id)
	fmt.Printf("Date eintrag1 lautet: %v\n", eintrag1.date)

	eintrag2 := NewEntry(2)
	fmt.Printf("Id eintrag2 lautet: %v\n", eintrag2.id)
	fmt.Printf("Date eintrag2 lautet: %v\n", eintrag2.date)

	var jsonData []byte
	jsonData, err := json.Marshal(eintrag1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))

}
