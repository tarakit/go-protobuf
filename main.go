package main

import (
	"fmt"

	"github.com/tarakit/go-protobuf/src/person"

)

func main()  {
	doSomething()
}


func doSomething()  {
	ps := simple.protobuf.person{
		Name : "Kit Tara",
		Age : 23,
		Gender : "Male",
		PhoneNumber : []int32{069320730, 099332232}
	}

	fmt.Println(ps)
}