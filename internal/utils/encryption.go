package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"io"
)

func DecryptData(cipherText string) string {
	cipherTextByte, err := hex.DecodeString(cipherText)
	if err != nil {
		Logger.Error("failed to convert ciphertext to byte array", zap.Error(err))
	}
	block, err := aes.NewCipher(hashMachineId(getMachineId()))
	if err != nil {
		Logger.Error("failed to generate cipher block", zap.Error(err))
	}
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		Logger.Error("failed to wrap block", zap.Error(err))
	}
	nonceSize := aesGcm.NonceSize()
	nonce, encryptedText := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]
	plaintext, err := aesGcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		Logger.Error("failed to decrypt ciphertext", zap.Error(err))
	}
	return string(plaintext)
}

func EncryptData(secret string) string {
	plaintext := []byte(secret)
	block, err := aes.NewCipher(hashMachineId(getMachineId()))
	if err != nil {
		Logger.Error("failed to generate cipher block", zap.Error(err))
	}
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		Logger.Error("failed to wrap cipher block")
	}
	nonce := make([]byte, aesGcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		Logger.Error("failed to generate nonce", zap.Error(err))
	}
	cipherText := aesGcm.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", cipherText)
}

func hashMachineId(machineId string) []byte {
	hash := sha256.New()
	hash.Write([]byte(machineId))
	computedHash := hash.Sum(nil)
	return computedHash
}
