package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
)

func main() {

	fmt.Println("RPC Server started...")

	// In general ways, should be 'package.NewCollege()'
	mit := NewCollege()

	rpc.Register(mit)
	rpc.HandleHTTP()
	// rpc.HandleHTTP(rpcPath, debugPath string)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "RPC Server Live!")
		},
	)

	http.ListenAndServe(":9093", nil)

}
