package main

import (
	"hash"
	"io"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/rand"
)

type MessageDecryptor struct {
	Hash       hash.Hash
	Reader     io.Reader
	PrivateKey rsa.PrivateKey
	label      []byte
}

func (md MessageDecryptor) DecryptMessage(message []byte) ([]byte, error) {
	decryptedMessage, error := rsa.DecryptOAEP(md.Hash, md.Reader, &md.PrivateKey, message, md.label)
	if error != nil {
		return nil, error
	}
	return decryptedMessage, nil
}

func NewMessageDecryptor(PrivatKey rsa.PrivateKey, label []byte) (*MessageDecryptor) {
	return &MessageDecryptor{
		Hash:       sha256.New(),
		PrivateKey: PrivatKey,
		Reader:     rand.Reader,
		label:      label,
	}
}
