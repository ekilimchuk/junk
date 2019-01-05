package main

import (
	"../util"
	"crypto/rsa"
)

type Keys struct {
	ServerPublicKey  *rsa.PublicKey
	AgentToken      string
	AgentPrivateKey *rsa.PrivateKey
	AgentPublicKey  *rsa.PublicKey
}

func (k *Keys) LoadKeys(c util.AgentConfig) error {
	ServerPublicKey, err := util.LoadPublicKey(c.ServerPublicKey)
	if err != nil {
		return err
	}
	AgentToken, err := util.LoadAgentToken(c.AgentToken)
	if err != nil {
		return err
	}
	AgentPrivateKey, err := util.LoadPrivateKey(c.AgentPrivateKey)
	if err != nil {
		return err
	}
	AgentPublicKey, err := util.LoadPublicKey(c.AgentPublicKey)
	if err != nil {
		return err
	}
	k.ServerPublicKey = ServerPublicKey
	k.AgentToken = AgentToken
	k.AgentPrivateKey = AgentPrivateKey
	k.AgentPublicKey = AgentPublicKey
	return nil
}
