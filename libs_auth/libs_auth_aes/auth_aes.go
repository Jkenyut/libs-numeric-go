package libs_auth_aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

// AESCustomRequest used for input into EncryptMessage for data encryption.
type AESCustomRequest struct {
	Request any `json:"request,omitempty"`
}

// AESCustomResponse used for input into DecryptMessage for data encryption.
type AESCustomResponse struct {
	Response any `json:"response,omitempty"`
}

// EncryptMessage data encryption via AESCustomRequest input using the AES algorithm AesEncrypt
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

// DecryptMessage data description via AESCustomResponse input using the AES algorithm
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

// AESEncrypt the function called to perform encryption
func AESEncrypt(key []byte, data any) (encode string, err error) {
	jsonData, err := json.Marshal(AESCustomRequest{Request: data})
	if err != nil {
		return encode, err
	}
	encryptedData, err := EncryptMessage(key, jsonData)
	if err != nil {
		return encode, err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// AESDecrypt the function called to perform description
func AESDecrypt(key []byte, plaintext string) (out *AESCustomRequest, err error) {
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
