package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenSign(accessKey, secretKey string, timestamp int64) (string, int64, error) {

	// 2. 拼接字符串 {access_key}_{timestamp}
	message := fmt.Sprintf("%s_%d", accessKey, timestamp)

	// 3. 使用HMAC-SHA256计算哈希
	h := hmac.New(sha256.New, []byte(secretKey))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", 0, err
	}

	// 4. 转换为十六进制大写字符串
	signature := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(signature), timestamp, nil
}

func VerifySign(accessKey, secretKey string, timestamp int64, rawSign string) bool {
	// 1. 获取当前UTC时间戳（毫秒）

	sign, _, _ := GenSign(accessKey, secretKey, timestamp)
	return sign == rawSign
}
