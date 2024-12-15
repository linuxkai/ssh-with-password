package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"net"
	"strconv"
)

func ValidateIP(ip string) (bool, error) {
	if net.ParseIP(ip) == nil {
		return false, errors.New("invalid IP address.")
	}
	return true, nil
}

func ValidatePort(portStr string) (bool, error) {
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return false, errors.New("invalid port number: must be 1-65535")
	}
	return true, nil
}

func ValidateNum(numStr string) (bool, error) {
	_, err := strconv.Atoi(numStr)
	if err != nil {
		return false, errors.New("invalid number: must be an integer")
	}

	return true, nil
}

var key = "LinuxKaiLinuxKai"

// 字符串对称加密
func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 初始化向量
	iv := []byte(key)[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)

	plaintextBytes := []byte(plainText)
	ciphertext := make([]byte, len(plaintextBytes))
	stream.XORKeyStream(ciphertext, plaintextBytes)

	// 使用 Base64 编码结果
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(cipherText string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ciphertext, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	iv := []byte(key)[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
