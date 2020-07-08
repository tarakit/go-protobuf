package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	enumPb "github.com/protobuf-testing/src/enum"
	simplepb "github.com/protobuf-testing/src/person"
)

func main() {
	ps := doSomething()
	// readAndWriteFileDemo(ps)

	psAsString := toJSON(ps)
	fmt.Println("Protobuf msg to JSON" + psAsString)

	ps2 := &simplepb.Person{}
	fromJSON(psAsString, ps2)
	fmt.Println("JSON to Protobuf msg:", ps2)

	// doEnumTest()
}

func doEnumTest() {
	enum := enumPb.EnumDay{
		Id:           1,
		DayOfTheWeek: enumPb.DayofTheWeek_SATURDAY,
	}
	fmt.Println(enum)
}

func fromJSON(in string, pb proto.Message) *proto.Message {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		fmt.Println("can't unmarshling json to protobuf")
		log.Fatal(err)
	}
	return &pb
}

func toJSON(pb proto.Message) string {
	marshler := jsonpb.Marshaler{}
	out, err := marshler.MarshalToString(pb)
	if err != nil {
		println("can't convert to JSON")
		log.Fatal(err)
	}
	return out
}

func readAndWriteFileDemo(pb proto.Message) {
	// writeToFile("simple.bin", ps)

	// var ps2 *simplepb.Person
	ps2 := &simplepb.Person{}
	readFromFile("simple.bin", ps2)
	fmt.Println(ps2)
}

func readFromFile(fileName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		println("can't put the byte into the proto struct here")
		return err
	}
	fmt.Println(pb)

	return nil
}

func writeToFile(fileName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("File has been written")
	return nil
}

func doSomething() *simplepb.Person {

	ps := simplepb.Person{
		Name:        "Kit Tara",
		Age:         23,
		Gender:      "Male",
		PhoneNumber: []string{"06, 23, 25", "02, 25, 45"},
	}

	// fmt.Println(ps)
	// ps.Name = "Toro Gaga"
	// fmt.Println(ps)

	return &ps
}
