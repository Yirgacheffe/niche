package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	fmt.Println("os.Args => ", os.Args)

	fmt.Printf("os.PathSeparator => %v\n", os.PathSeparator)
	fmt.Printf("os.PathSeparator => %c\n", os.PathSeparator)

	fmt.Printf("os.DevNull => %v\n", os.DevNull)

	// os env
	envs := os.Environ()
	fmt.Printf("os.Environ => %v\n", envs)

	fruit := os.Getenv("FRUIT")
	fmt.Printf("os.Getenv(\"FRUIT\") => %s\n", fruit)

	orange, oExists := os.LookupEnv("ORANGE")
	fmt.Printf("os.LookupEnv(\"ORANGE\") => %v (exists: %v)\n", orange, oExists)

	err := os.Setenv("FRUIT", "apple")
	if err != nil {
		fmt.Println(err)
	}

	apple := os.Getenv("FRUIT")
	fmt.Printf("os.Getenv(\"FRUIT\") => %s\n", apple)

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println(err)
	}

	apple = os.Getenv("FRUIT")
	fmt.Printf("os.Getenv(\"FRUIT\") => %s\n", apple)

	// os expand, it should be string functions ???
	// Replace the place holder
	expdMapper := func(ph string) string {
		switch ph {
		case "FRUIT":
			return "mongo"
		case "CAR":
			return "VW"
		default:
			return "<empty>"
		}
	}

	raw := "I am eating $FRUIT and driving ${CAR}"
	formatted := os.Expand(raw, expdMapper)

	fmt.Println(formatted)

	// os.ExpandEnv
	os.Setenv("COUNTRY", "Japan")
	raww := "I am living in ${COUNTRY}."
	fmt.Printf("os.ExpandEnv => %s\n", os.ExpandEnv(raww))

}
