package app

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

func createHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	var err error

	key, err := hex.DecodeString(createHash(passphrase))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	var err error
	key, err := hex.DecodeString(createHash(passphrase))
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func EncryptFile(filename string, passphrase string) error {
	var err error

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	values, err := encrypt(data, passphrase)
	if err != nil {
		return err
	}

	os.Remove(filename)

	f, err := os.Create(filename + ".d")
	if err != nil {
		return err
	}

	defer f.Close()

	_, _ = f.Write(values)
	return nil
}

func DecryptFile(filename string, passphrase string) error {
	var err error
	data, err := ioutil.ReadFile(filename + ".d")
	if err != nil {
		return err
	}

	values, err := decrypt(data, passphrase)
	if err != nil {
		return err
	}

	os.Remove(filename + ".d")

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	_, _ = f.Write(values)
	return nil
}
