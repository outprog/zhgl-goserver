package md5passwd

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Get(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hasher.Sum(nil)))
}
