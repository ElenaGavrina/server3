package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func DecryptMessage(message string, key []byte) (string, error) {

	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "could not base64 decode", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "could not create new cipher", err
	}

	if len(cipherText) < aes.BlockSize {
		return "invalid ciphertext block size", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}