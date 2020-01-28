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

type EntryFactory struct{}

func (ef *EntryFactory) NewEntry(id int) *entry {

	return &entry{
		Id:   id,
		Date: time.Now(),
	}
}

func (e *entry) DoSerialize() {
	var jsonData []byte
	jsonData, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonData)
}

func (e *entry) DoDeSerialize() *entry {
	return nil
}
