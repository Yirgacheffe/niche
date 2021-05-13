package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Global constant and variables
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

	// System and User info
	hostName, err := os.Hostname()
	if err == nil {
		fmt.Printf("os.HostName => %s\n", hostName)
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		fmt.Printf("os.UserHomeDir => %s\n", homeDir)
	}

	// current work directory and change it
	currWd, _ := os.Getwd()
	fmt.Printf("os.Getwd [before]=> %v\n", currWd)

	os.Chdir("..")
	aftrWd, _ := os.Getwd()
	fmt.Printf("os.Getwd [after ]=> %v\n", aftrWd)

	os.Chdir("/Users/aaron/proj/golang-exec/niche/prac")
	aftrrWd, _ := os.Getwd()
	fmt.Printf("os.Getwd [after2]=> %v\n", aftrrWd)

	// executable program path
	execDir, _ := os.Executable()
	fmt.Printf("os.Executable => %v\n", execDir)

	tmpDir := os.TempDir()
	fmt.Printf("os.TempDir => %v\n", tmpDir)

}
