package main

import (
	"io"
	"crypto/rsa"
)

type PublicKeyGenerator interface {
	GeneratePublicKey() (rsa.PublicKey, error)
}
type KeyGenerator struct {
	Reader  io.Reader
	BitSize int
}

func NewKeyGenerator(r io.Reader, size int) (*KeyGenerator) {
	return &KeyGenerator{
		Reader:  r,
		BitSize: size,
	}
}
func (kg KeyGenerator) GenerateKey() (*rsa.PrivateKey, error) {
	key, error := rsa.GenerateKey(kg.Reader, kg.BitSize)
	if error != nil {
		return nil, error
	}
	return key, nil
}
