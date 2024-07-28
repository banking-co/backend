package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
)

func VkValidate(data map[string][]string, secret string) (bool, error) {
	var hashSecret []byte
	var hash = data["sign"][0]

	var buffer bytes.Buffer

	var keys []string
	for key := range data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		if key != "hash" {
			buffer.WriteString(key)
			buffer.WriteString(fmt.Sprintf("%v", data[key]))
		}
	}

	if hashSecret == nil {
		hash := sha256.New()
		hash.Write([]byte(secret))
		hashSecret = hash.Sum(nil)
	}

	impHmac := hmac.New(sha256.New, hashSecret)
	impHmac.Write(buffer.Bytes())

	return hex.EncodeToString(impHmac.Sum(nil)) == hash, nil

}
