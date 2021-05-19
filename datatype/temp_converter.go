package main

import (
	"fmt"
	"os"
	"strconv"
)

type (
	Celsius    float64
	Fahrenheit float64
)

const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func (f Fahrenheit) String() string { return fmt.Sprintf("%g", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%g", c) }

//
func celsius2fahrenheit(t float64) float64 { return 9.0/5.0*t + 32 }
func fahrenheit2celsius(t float64) float64 { return (t - 32) * 5.0 / 9.0 }

func fToc(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func usage() {
	fmt.Println("Usage: temperature_converter <mode> <temperature>")
	fmt.Println()
	fmt.Println("This program converts temperatures between Celsius and Fahrenheit")
	fmt.Println("'mode' is either 'c2f' or 'f2c'")
	fmt.Println("'temperature' is a floating point number to be converted according to mode")

	os.Exit(1)
}

func main() {

	if len(os.Args) != 3 {
		usage()
	}

	mode := os.Args[1]
	if mode != "c2f" && mode != "f2c" {
		usage()
	}

	t, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		usage()
	}

	var converted float64

	// Protected programming style
	if mode == "f2c" {
		converted = fahrenheit2celsius(t)
	} else if mode == "c2f" {
		converted = celsius2fahrenheit(t)
	} else {
		usage()
	}

	fmt.Println(converted)

	var c Celsius
	var f Fahrenheit

	// fmt.Println(c == f)
	fmt.Println(c >= 0)
	fmt.Println(c == fToc(f))

}
