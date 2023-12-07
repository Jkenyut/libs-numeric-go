package auth_aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

// aes custom data

type AESCustom struct {
	Data any
}

func EncryptMessage(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func DecryptMessage(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func AesEncrypt(key []byte, data any) (encode string, err error) {
	jsonData, err := json.Marshal(AESCustom{Data: data})
	if err != nil {
		return encode, err
	}
	encryptedData, err := EncryptMessage(key, jsonData)
	if err != nil {
		return encode, err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

func AesDecrypt(key []byte, plaintext string) (out *AESCustom, err error) {
	decodeString, err := base64.StdEncoding.DecodeString(plaintext)
	if err != nil {
		return nil, err
	}
	DecryptedData, err := DecryptMessage(key, decodeString)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(DecryptedData, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
