package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// CalculateContentHash 计算文本内容的SHA-256特征值
func CalculateContentHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}
