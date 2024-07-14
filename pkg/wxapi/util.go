package wxapi

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func CheckSignature(token, timestamp, nonce, signature string) bool {
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := strings.Join(tmpArr, "")
	hasher := sha1.New()
	hasher.Write([]byte(tmpStr))
	hexStr := hex.EncodeToString(hasher.Sum(nil))
	return hexStr == signature
}
