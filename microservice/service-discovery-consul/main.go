package main

import "discovery/client"

func main() {
	if err := client.Exec(); err != nil {
		panic(err)
	}
}
