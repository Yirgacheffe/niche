package main

import (
	"flag"
	"fmt"
)

// Config is the holder for our flags
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup initilize a config flag from the passed in
func (c *Config) Setup() {

	// you can set a flag directly like so: var someVar = flag.String("flag_name", "default_val", "description")

	// but in practice putting it in a struct is generally
	// better longhand
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// shorthand
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10")

	// custom variable type
	flag.Var(&c.countTheWays, "c", "comma seperate list of integers")

}

// GetMessage use all the config variants then return a sentence
func (c *Config) GetMessage() string {
	msg := c.subject

	if !c.isAwesome {
		msg += " is NOT awesome"
	} else {
		msg += " is awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}
