package client

import (
	"github.com/hashicorp/consul/api"
)

type Client interface {
	Register(tags []string) error
	Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error)
}

type client struct {
	client  *api.Client
	name    string
	address string
	port    int
}

func NewClient(cfg *api.Config, name, address string, port int) (Client, error) {
	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	cli := &client{client: c, name: name, address: address, port: port}
	return cli, nil
}

func (c *client) Register(tags []string) error {
	reg := &api.AgentServiceRegistration{
		ID:      c.name,
		Name:    c.name,
		Port:    c.port,
		Address: c.address,
		Tags:    tags,
	}

	return c.client.Agent().ServiceRegister(reg)
}

func (c *client) Service(service, tag string) ([]*api.ServiceEntry, *api.QueryMeta, error) {
	return c.client.Health().Service(service, tag, false, nil)
}
