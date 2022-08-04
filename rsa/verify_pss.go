package rsa

import (
	"crypto/rsa"
	"encoding/base64"
)

func (r *FastRSA) VerifyPSS(signature, message, hashName, saltLengthName, publicKey string) (bool, error) {
	messsageBytes, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return false, err
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	return r.verifyPSS(signatureBytes, messsageBytes, hashName, saltLengthName, publicKey)
}

func (r *FastRSA) VerifyPSSBytes(signature, message []byte, hashName, saltLengthName, publicKey string) (bool, error) {
	return r.verifyPSS(signature, message, hashName, saltLengthName, publicKey)
}

func (r *FastRSA) verifyPSS(signature, message []byte, hashName, saltLengthName, publicKey string) (bool, error) {
	public, err := r.readPublicKey(publicKey)
	if err != nil {
		return false, err
	}

	hash := getHashInstance(hashName)
	_, err = hash.Write(message)
	if err != nil {
		return false, err
	}

	err = rsa.VerifyPSS(public, getHashType(hashName), hash.Sum(nil), signature, &rsa.PSSOptions{
		SaltLength: getSaltLength(saltLengthName),
		Hash:       getHashType(hashName),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
