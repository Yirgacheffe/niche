package client

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

const (
	address = "localhost:8500"
	name    = "discovery"
)

func Exec() error {
	config := consul.DefaultConfig()
	config.Address = address

	cli, err := NewClient(config, name, "localhost", 8080)
	if err != nil {
		return err
	}
	if err := cli.Register([]string{"Go", "Awesome"}); err != nil {
		return err
	}

	entries, _, err := cli.Service(name, "Go")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("%#v\n", entry.Service)
	}

	return nil
}
