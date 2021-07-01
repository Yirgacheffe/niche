package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"
const usage = `Usage: 
%s [command] 
Commands: 
	Greet 
	Version
`

const greetUsage = `Usage:
%s greet name [Flag]
Positional Arguments:
name
	the name to greet
Flags:
`

// MenuConf holds all the levels for a nested cmd line argument
type MenuConf struct {
	Goodbye bool
}

// SetupMenu initializes the base flags
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu return a flag set for a submenu
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	subMenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	subMenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")

	subMenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		subMenu.PrintDefaults()
	}
	return subMenu
}

// Greet will be invoked by the greet command
func (m *MenuConf) Greet(name string) {
	if !m.Goodbye {
		fmt.Printf("Hello, %s !\n", name)
	} else {
		fmt.Printf("Goodbye %s !\n", name)
	}
}

// Version prints the current version that is stored as a const
func (m *MenuConf) Version() {
	fmt.Println("Version: ", version)
}

func main() {
	c := MenuConf{}
	menu := c.SetupMenu()
	menu.Parse(os.Args[1:])

	// we use arguments to switch between commands flags
	// are also an argument
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			f.Parse(os.Args[3:])
			c.Greet(os.Args[2])
		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}

	} else {
		menu.Usage()
		return
	}
}
