package main

import "errors"

const (
	OK = "+OK\r\n\n"
)

type Server interface {
	Lock()
	Unlock()
	GetData(key string) interface{}
	SetData(key string, value interface{})
}

type plugin struct {
	name        string
	commands    []string
	description string
}

var Plugin plugin

func (p *plugin) Name() string {
	return p.name
}

func (p *plugin) Commands() []string {
	return p.commands
}

func (p *plugin) Description() string {
	return p.description
}

func (p *plugin) HandleCommand(cmd []string, server interface{}) ([]byte, error) {
	switch len(cmd) {
	default:
		return nil, errors.New("wrong number of arguments for PING command")
	case 1:
		return []byte("+PONG\r\n\n"), nil
	case 2:
		return []byte("+" + cmd[1] + "\r\n\n"), nil
	}
}

func init() {
	Plugin.name = "PingCommand"
	Plugin.commands = []string{"ping"}
	Plugin.description = "Handle PING command"
}
