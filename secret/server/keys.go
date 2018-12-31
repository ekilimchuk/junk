package main

import (
	"../util"
	"crypto/rsa"
)

type Keys struct {
	ServerPrivateKey *rsa.PrivateKey
	ServerPublicKey  *rsa.PublicKey
	ServerToken      string
	ClientPublicKey  *rsa.PublicKey
}

func (k *Keys) LoadKeys(c util.ServerConfig) error {
	ServerPrivateKey, err := util.LoadPrivateKey(c.ServerPrivateKey)
	if err != nil {
		return err
	}
	ServerPublicKey, err := util.LoadPublicKey(c.ServerPublicKey)
	if err != nil {
		return err
	}
	ServerToken, err := util.LoadClientToken(c.ServerToken)
	if err != nil {
		return err
	}
	ClientPublicKey, err := util.LoadPublicKey(c.ClientPublicKey)
	if err != nil {
		return err
	}
	k.ServerPrivateKey = ServerPrivateKey
	k.ServerPublicKey = ServerPublicKey
	k.ServerToken = ServerToken
	k.ClientPublicKey = ClientPublicKey
	return nil
}
