package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"crypto/rand"
)

func TestNewKeyGenerator(t *testing.T) {
	kGen := NewKeyGenerator(rand.Reader, 234)
	assert.IsType(t, &KeyGenerator{}, kGen)
}

func TestKeyGeneratorGeneratePublicKey(t *testing.T) {
	kGen := NewKeyGenerator(rand.Reader, 234)
	key, err := kGen.GenerateKey()
	fmt.Println(key)
	assert.NoError(t, err)
	assert.NotEmpty(t, key)
}

func TestNewMessageEncryptor(t *testing.T) {
	kGEn := NewKeyGenerator(rand.Reader, 2340)
	key, err := kGEn.GenerateKey()
	assert.NoError(t, err)

	label := "This is a test label"
	message := "This is a test message"

	encyptor := NewMessageEncryptor(&key.PublicKey, []byte(label))
	encMessage, err := encyptor.EncryptMessage(message)
	assert.NoError(t, err)
	assert.NotEmpty(t, encMessage)
	assert.Equal(t, string(encyptor.Label), label)
}
func TestMessageDecryptorDecryptMessage(t *testing.T) {
	kGEn := NewKeyGenerator(rand.Reader, 2340)
	key, err := kGEn.GenerateKey()
	assert.NoError(t, err)

	label := "This is a test label"
	message := "This is a test message"

	encyptor := NewMessageEncryptor(&key.PublicKey, []byte(label))
	encMessage, err := encyptor.EncryptMessage(message)

	decryptor := NewMessageDecryptor(*key, []byte(label))
	decrypted, err := decryptor.DecryptMessage(encMessage)
	assert.NoError(t, err)
	assert.Equal(t, string(decrypted), message)
}
