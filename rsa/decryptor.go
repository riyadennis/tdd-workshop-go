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
}

func (md MessageDecryptor) DecryptMessage(message []byte) ([]byte, error) {
	decryptedMessage, error := rsa.DecryptPKCS1v15(md.Reader, &md.PrivateKey, message)
	if error != nil {
		return nil, error
	}
	return decryptedMessage, nil
}

func NewMessageDecryptor(PrivatKey rsa.PrivateKey) (*MessageDecryptor) {
	return &MessageDecryptor{
		Hash:       sha256.New(),
		PrivateKey: PrivatKey,
		Reader:     rand.Reader,
	}
}
