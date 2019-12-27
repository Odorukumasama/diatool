package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type entry struct {
	id   int
	date time.Time
}

func NewEntry(id int) (*entry, error) {
	if id < 1 {
		return nil, errors.New("ID should be > 0")
	}

	e := new(entry)
	e.id = id
	e.date = time.Now()

	return e, nil
}

func main() {

	eintrag1, error := NewEntry(1)
	if error != nil {
		fmt.Printf("Fehler: %s\n", error)
		return
	}
	fmt.Printf("Id eintrag1 lautet: %v\n", eintrag1.id)
	fmt.Printf("Date eintrag1 lautet: %v\n", eintrag1.date)

	eintrag2, error := NewEntry(2)
	if error != nil {
		fmt.Printf("Fehler: %s\n", error)
	}
	fmt.Printf("Id eintrag2 lautet: %v\n", eintrag2.id)
	fmt.Printf("Date eintrag2 lautet: %v\n", eintrag2.date)

	var jsonData []byte
	jsonData, err := json.Marshal(eintrag2.id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", jsonData)

}
