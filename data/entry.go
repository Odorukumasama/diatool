// entry
package entry

import (
	"encoding/json"
	"fmt"
	"time"
)

type entry struct {
	Id   int       `json:"id"`
	Date time.Time `json:"date"`
}

func NewEntry(id int) *entry {
	e := new(entry)
	e.Id = id

	e.Date = time.Now()

	return e
}

func DoSerialize(o *entry) {
	var jsonData []byte
	jsonData, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)
}
