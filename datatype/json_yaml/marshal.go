package main

import "fmt"

// MarshalAll takes
func MarshalAll() error {

	j := JSONData{
		Name: "Name2", Age: 30,
	}

	y := YAMLData{
		Name: "Name3", Age: 40,
	}

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}
	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}
	fmt.Println("YAML Marshal=", yamlRes.String())

	return nil

}
