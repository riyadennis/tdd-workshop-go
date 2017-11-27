package main

import (
	"hash"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/rand"
)

type Encryptor interface {
	EncryptMessage()
}

type MessageEncryptor struct {
	Hash      hash.Hash
	PublicKey *rsa.PublicKey
}

func NewMessageEncryptor(PublicKey *rsa.PublicKey) (*MessageEncryptor) {
	return &MessageEncryptor{
		Hash:      sha256.New(),
		PublicKey: PublicKey,
	}
}

func (me MessageEncryptor) EncryptMessage(message string) ([]byte, error) {
	enMessage, error := rsa.EncryptPKCS1v15(
		rand.Reader,
		me.PublicKey,
		[]byte(message),
	)
	if error != nil {
		return nil, error
	}
	return enMessage, nil
}
