package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CountTheWays is a custom type that will be read a flag
type CountTheWays []int

func (c *CountTheWays) String() string {
	result := ""

	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}

	return result
}

// Set will be used by the flag package
func (c *CountTheWays) Set(value string) error {

	values := strings.Split(value, ",")

	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}

	return nil // Re, return nil as no error occured

}
