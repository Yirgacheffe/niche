package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// DollarToPennies takes a dollar amount
// as a string, i.e. 1.00, 55.12 etc and converts it into an int64
func DollarToPennies(amount string) (int64, error) {

	// check if amount can be convert to valid float
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	groups := strings.Split(amount, ".")
	result := groups[0]

	// a base string for result composing
	cents := ""

	if len(groups) == 2 {
		cents = groups[1]
	}

	if len(cents) > 2 {
		return 0, errors.New("invalid cents format")
	}

	// pad with 0, this will be 2 0's if there was no .
	for len(cents) < 2 {
		cents += "0"
	}

	result += cents
	return strconv.ParseInt(result, 10, 64)

}

// PenniesToDollar takes a penny amount as an int64 and returns
// a dollar string representation
func PenniesToDollar(amount int64) string {

	result := strconv.FormatInt(amount, 10)
	isNeg := false

	if result[0] == '-' {
		result = result[1:]
		isNeg = true
	}

	// left pad with 0 if we're passed in value < 100
	for len(result) < 3 {
		result = "0" + result
	}

	l := len(result)
	result = result[0:l-2] + "." + result[l-2:]

	// left pad '-' if negnative as checked
	if isNeg {
		result = "-" + result
	}

	return result // Re, return the string value in dollars

}

func main() {

	inputAmt := "-15.93"
	fmt.Printf("input amount is %s dollars\n", inputAmt)

	pennies, err := DollarToPennies(inputAmt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("input amount convert to %d pennies\n", pennies)

	// add 15 cents into last value
	pennies += 15
	dollars := PenniesToDollar(pennies)

	fmt.Printf("add 15 cents, now the value is %s\n", dollars)

}
