package main

import (
	"flag"
	"fmt"
)

// Celsius and Fahrenheit temperature unit
type (
	Celsius    float64
	Fahrenheit float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%g", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g", f) }
func fToc(f Fahrenheit) Celsius     { return Celsius((f - 32) * 5 / 9) }

// celsiusFlag
type celsiusFlag struct{ Celsius }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var (
		unit string
		v    float64
	)

	fmt.Sscanf(s, "%f%c", &v, &unit) // ignore the error
	switch unit {
	case "C", "c":
		f.Celsius = Celsius(v)
		return nil
	case "F", "f":
		f.Celsius = fToc(Fahrenheit(v))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func main() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")

	flag.Parse()
	fmt.Println(*temp)
}
