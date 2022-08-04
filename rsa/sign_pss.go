package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func (r *FastRSA) SignPSS(message, hashName, saltLengthName, privateKey string) (string, error) {
	messsageBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", err
	}

	output, err := r.signPSS(messsageBytes, hashName, saltLengthName, privateKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(output), nil
}

func (r *FastRSA) SignPSSBytes(message []byte, hashName, saltLengthName, privateKey string) ([]byte, error) {
	return r.signPSS(message, hashName, saltLengthName, privateKey)
}

func (r *FastRSA) signPSS(message []byte, hashName, saltLengthName, privateKey string) ([]byte, error) {

	private, err := r.readPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	hash := getHashInstance(hashName)
	_, err = hash.Write(message)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPSS(rand.Reader, private, getHashType(hashName), hash.Sum(nil), &rsa.PSSOptions{
		SaltLength: getSaltLength(saltLengthName),
		Hash:       getHashType(hashName),
	})
	if err != nil {
		return nil, err
	}

	return signature, nil
}
