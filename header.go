package go_mtpay

import (
	"github.com/spf13/cast"
)

func getHeaders(accessKey, sign string, unixMilli int64) map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"charset":      "utf-8",
		"timestamp":    cast.ToString(unixMilli), //time.Now().UTC().UnixMilli()),
		"access_key":   accessKey,
		"signature":    sign,
	}
}
