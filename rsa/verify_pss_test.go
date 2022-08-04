package rsa

import (
	"encoding/base64"
	"testing"
)

func TestFastRSA_VerifyPSS(t *testing.T) {

	instance := NewFastRSA()
	inputMessageB64 := base64.StdEncoding.EncodeToString([]byte(inputMessage))
	output, err := instance.VerifyPSS(signedPSS, inputMessageB64, "sha512", "equalsHash", publicKey)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("output:", output)
}
