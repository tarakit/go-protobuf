package main

import (
	"fmt"

	simplepb "github.com/protobuf-testing/src"
)

func main() {
	doSomething()
}

func doSomething() *simplepb.Person {

	ps := simplepb.Person{
		Name:        "Kit Tara",
		Age:         23,
		Gender:      "Male",
		PhoneNumber: []string{"06, 23, 25", "02, 25, 45"},
	}

	fmt.Println(ps)
	// ps.Name = "Toro Gaga"
	// fmt.Println(ps)

	return &ps
}
