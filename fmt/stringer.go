package main

import (
	"fmt"
)

type animal struct {
	name string
}

func (a *animal) String() string {
	return fmt.Sprintf("I'm %s", a.name)
}

func main() {
	cat := animal{"cat"}
	fmt.Printf("%s\n", cat)
	fmt.Printf("%s\n", &cat)
	fmt.Printf("%s\n", cat.String())
}
