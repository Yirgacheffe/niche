package main

import (
	"log"
	"net/http"
	"raft/consensus"
)

func main() {
	consensus.Config(3)
	http.HandleFunc("/", consensus.Handler)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
