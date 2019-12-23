package main

import (
	//"encoding/json"
	"fmt"
	//"time"
)

type entry struct {
	id int
	//date time.Time
}

func NewEntry() *entry {
	e := new(entry)
	e.id = 1

	return e
}

func main() {
	eintrag1 := NewEntry()
	fmt.Printf("Id lautet: %v\n", eintrag1.id)
}
