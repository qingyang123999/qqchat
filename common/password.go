package common

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// ============================     hash 加密       =======================================

// GeneratePasswordHash   生成 密码hash
// param  password 需加密的明文密码
// param  cost 工作因子(cost) cost‌：计算成本(4-31)，决定哈希迭代次数。建议值： bcrypt.MinCost(4)：测试环境使用 bcrypt.DefaultCost(10)：生产环境默认值10
// bcrypt是一种基于Blowfish加密算法的密码哈希函数，通过盐值(salt)和工作因子(cost)增强安全性
func GeneratePasswordHash(password string, cost int) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost) // 当填写的cost小于4时。默认为10
	return string(bytes), err
}

// VerifyPasswordHash 校验 密码hash
// param  hashedPassword：先前存储的哈希密码
// param  password：待验证的明文密码
func VerifyPasswordHash(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// ============================     AES 加密     ECB 模式加密   PKCS7 填充  =======================================

// PKCS7Padding pads an input slice to be a multiple of the block size, using PKCS#7 padding.
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding removes the PKCS#7 padding from an input slice.
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unpadding := int(origData[length-1])
	if unpadding > length {
		return nil, errors.New("unpadding size exceeded")
	}
	return origData[:(length - unpadding)], nil
}

// BytesToHex 将字节切片转换为十六进制字符串
func BytesToHex(data []byte) string {
	hexString := ""
	for _, b := range data {
		// 使用格式化字符串将每个字节转换为两位十六进制数
		hexString += fmt.Sprintf("%02x", b)
	}
	return hexString
}

// Encrypt encrypts plaintext using AES encryption in ECB mode with PKCS7 padding.
func AesEcbPkcs7Encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// PKCS7 填充
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	if len(plaintext)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}

	// ECB 模式加密
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], plaintext[i:i+blockSize])
	}

	return string(ciphertext), nil
}

// Decrypt decrypts ciphertext using AES decryption in ECB mode with PKCS7 padding.
func AesEcbPkcs7Decrypt(ciphertextHex string, key []byte) (string, error) {
	ciphertext := []byte(ciphertextHex)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	// ECB 模式解密
	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}

	// 去除 PKCS7 填充
	plaintext, err = PKCS7UnPadding(plaintext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// ============================     md5 加密       =======================================

// Md5WithSalt 生成带盐值的MD5哈希
// originalText: 原始字符串 待加密字符串
// salt: 盐值字符串
func Md5WithSalt(originalText string, salt string) (md5Password string) {
	hasher := md5.New()                        // 创建MD5哈希对象
	hasher.Write([]byte(originalText + salt))  // 拼接原始文本和盐值  并 写入要计算的数据
	return hex.EncodeToString(hasher.Sum(nil)) // 将哈希值转换为16进制字符串
}
