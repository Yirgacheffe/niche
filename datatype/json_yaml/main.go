package main

func main() {

	if err := UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := MarshalAll(); err != nil {
		panic(err)
	}

	if err := JSONExampleWithMap(); err != nil {
		panic(err)
	}

}
