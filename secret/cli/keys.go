package main

import (
	"../util"
	"crypto/rsa"
)

type Keys struct {
	ServerPublicKey  *rsa.PublicKey
	ClientToken      string
	ClientPrivateKey *rsa.PrivateKey
	ClientPublicKey  *rsa.PublicKey
}

func (k *Keys) LoadKeys(c util.CliConfig) error {
	ServerPublicKey, err := util.LoadPublicKey(c.ServerPublicKey)
	if err != nil {
		return err
	}
	ClientToken, err := util.LoadClientToken(c.ClientToken)
	if err != nil {
		return err
	}
	ClientPrivateKey, err := util.LoadPrivateKey(c.ClientPrivateKey)
	if err != nil {
		return err
	}
	ClientPublicKey, err := util.LoadPublicKey(c.ClientPublicKey)
	if err != nil {
		return err
	}
	k.ServerPublicKey = ServerPublicKey
	k.ClientToken = ClientToken
	k.ClientPrivateKey = ClientPrivateKey
	k.ClientPublicKey = ClientPublicKey
	return nil
}
