// main
package main

import (
	entry "diatool/data"
	"fmt"
)

func main() {
	factory := entry.EntryFactory{}

	eintrag1 := factory.NewEntry(1)
	fmt.Printf("Id eintrag1 lautet: %v\n", eintrag1.Id)
	fmt.Printf("Date eintrag1 lautet: %v\n", eintrag1.Date)

	eintrag1.DoSerialize()

	eintrag2 := factory.NewEntry(2)
	fmt.Printf("Id eintrag2 lautet: %v\n", eintrag2.Id)
	fmt.Printf("Date eintrag2 lautet: %v\n", eintrag2.Date)

	eintrag2.DoSerialize()

}
