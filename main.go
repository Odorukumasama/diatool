// main
package main

import (
	"diatool/data"
	"fmt"
)

func main() {
	eintrag1 := entry.NewEntry(1)
	fmt.Printf("Id eintrag1 lautet: %v\n", eintrag1.Id)
	fmt.Printf("Date eintrag1 lautet: %v\n", eintrag1.Date)

	entry.DoSerialize(eintrag1)

	eintrag2 := entry.NewEntry(2)
	fmt.Printf("Id eintrag2 lautet: %v\n", eintrag2.Id)
	fmt.Printf("Date eintrag2 lautet: %v\n", eintrag2.Date)

	entry.DoSerialize(eintrag2)

}
