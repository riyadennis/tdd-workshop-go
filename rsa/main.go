package main

import (
	"crypto/rand"
	"github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	kg := NewKeyGenerator(rand.Reader, 2340)
	key, err := kg.GenerateKey()
	if err != nil {
		logrus.Fatalf("Unable to generate the key %s", err.Error())
	}
	message := "apikey:jjje:hh"

	messgeEnc := NewMessageEncryptor(&key.PublicKey)
	em, err := messgeEnc.EncryptMessage(message)
	if err != nil {
		logrus.Fatalf("Unable to generate the key %s", err.Error())
	}

	decrp := NewMessageDecryptor(*key)
	dm, err := decrp.DecryptMessage(em)
	if err != nil {
		logrus.Fatalf("Unable to generate the key %s", err.Error())
	}
	fmt.Println(string(dm))
}





