package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"flag"
	"os"
)

var (
	private = flag.String("private", "", "is a path for save a private pem file.")
	public  = flag.String("public", "", "is a path for save a public pem file.")
)

func saveKey(fileName string, key *rsa.PrivateKey) {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var k = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	err = pem.Encode(f, k)
	if err != nil {
		panic(err)
	}
}

func savePublicKey(fileName string, key rsa.PublicKey) {
	bytes, err := asn1.Marshal(key)
	if err != nil {
		panic(err)
	}
	var k = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = pem.Encode(f, k)
	if err != nil {
		panic(err)
	}
}
func main() {
	flag.Parse()
	if *private == "" || *public == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic(err)
	}
	saveKey(*private, key)
	savePublicKey(*public, key.PublicKey)
}
