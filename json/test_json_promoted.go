package main

import (
	"encoding/json"
	"fmt"
)

type POD struct {
	UserName  string
	Followers int
}

type XYZ struct {
	IsMale bool
	Email  string
}

type student struct {
	FirstName string
	LastName  string
	IsMale    bool
	POD
	XYZ
}

func main() {

	// "IsMale" and "Email" is not shown in parent fields
	// So will not promoted
	data := []byte(`
	{
		"FirstName": "hiluzan",
		"LastName": "salutobi",
		"IsMale": true,
		"UserName": "hisaludoe834",
		"Followers": 1982,
		"XYZ": { "IsMale": true, "Email": "hisalu@deo.com" }
	}`)

	var salu student

	fmt.Printf("Error: %v\n", json.Unmarshal(data, &salu))
	fmt.Printf("%#v\n", salu)

}
