package main

import (
	"encoding/xml"
	"fmt"
)

// Examples in this repository can run when the standard xml import is overridden with something like "../xml"

type Person struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
}

// Any change in the local encoding/xml package can be evaludated.
func main() {
	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}

	output, err := xml.MarshalIndent(v, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(string(output))
}
