package util

import (
	"encoding/json"
	"io/ioutil"
)

type Config interface {
	load(filePath string) error
}

type ServerConfig struct {
	ServerListen     string
	ServerPrivateKey string
	ServerPublicKey  string
	ServerToken      string
	ClientPublicKey  string
}

type CliConfig struct {
	Server           string
	ServerPublicKey  string
	ClientToken      string
	ClientPrivateKey string
	ClientPublicKey  string
}

type AgentConfig struct {
	Server          string
	ServerPublicKey string
	AgentToken      string
	AgentPrivateKey string
	AgentPublicKey  string
}

func (c *ServerConfig) load(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(bytes), c); err != nil {
		return err
	}
	return nil
}

func (c *CliConfig) load(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(bytes), c); err != nil {
		return err
	}
	return nil
}

func (c *AgentConfig) load(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(bytes), c); err != nil {
		return err
	}
	return nil
}

func LoadConfig(c Config, filePath string) error {
	return c.load(filePath)
}
