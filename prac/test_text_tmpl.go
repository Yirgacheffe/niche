package main

import (
	"fmt"
	"html/template"
	"os"
)

type Entry struct {
	Number int
	Square int
}

func main() {

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need the template file!")
		return
	}

	textTmpl := arguments[1]
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}
	var entries []Entry

	for _, i := range DATA {
		if len(i) == 2 {
			temp := Entry{Number: i[0], Square: i[1]}
			entries = append(entries, temp)
		}
	}

	t := template.Must(template.ParseGlob(textTmpl))
	t.Execute(os.Stdout, entries)

}
