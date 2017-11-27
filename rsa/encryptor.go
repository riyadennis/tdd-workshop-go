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
	Label     []byte
}

func NewMessageEncryptor(PublicKey *rsa.PublicKey, label []byte) (*MessageEncryptor) {
	return &MessageEncryptor{
		Hash:      sha256.New(),
		PublicKey: PublicKey,
		Label:     label,
	}
}

func (me MessageEncryptor) EncryptMessage(message string) ([]byte, error) {
	enMessage, error := rsa.EncryptOAEP(
		me.Hash,
		rand.Reader,
		me.PublicKey,
		[]byte(message),
		me.Label,
	)
	if error != nil {
		return nil, error
	}
	return enMessage, nil
}
