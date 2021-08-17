package main

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
)

type Client interface {
	Service(string, string) ([]*consul.ServiceEntry, *consul.QueryMeta, error)
	Register(string, string, string, int, []string) error
	DeRegister(string) error
}

type client struct {
	consul *consul.Client
}

func NewConsulClient(addr string) (*client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{
		consul: c,
	}, nil
}

// Register a service with consul local agent
func (c *client) Register(id string,name string, address string , port int, tags []string) error {
	reg := &consul.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Port:    port,
		Address: address,
		Tags: tags,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}
