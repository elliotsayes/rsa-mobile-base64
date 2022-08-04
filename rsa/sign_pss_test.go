package rsa

import (
	"encoding/base64"
	"testing"
)

func TestFastRSA_SignPSS(t *testing.T) {

	instance := NewFastRSA()

	inputMessageB64 := base64.StdEncoding.EncodeToString([]byte(inputMessage))
	output, err := instance.SignPSS(inputMessageB64, "sha512", "equalsHash", privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}

func TestFastRSA_SignPSSBigMessage(t *testing.T) {

	instance := NewFastRSA()
	inputMessage := getBigInputMessage()
	inputMessageB64 := base64.StdEncoding.EncodeToString([]byte(inputMessage))
	output, err := instance.SignPSS(inputMessageB64, "sha512", "equalsHash", privateKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}
