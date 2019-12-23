package main

import (
	//"encoding/json"
	"fmt"
	"time"
)

type entry struct {
	id   int
	date time.Time
}

func (e *entry) NewEntry(id int, date time.Time) *entry {
	e1 := new(entry)
	e1.id = 1

	return &e1{id}
}

func main() {

}
