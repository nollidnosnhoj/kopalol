package utils

import (
	"crypto/md5"
	"fmt"
)

func EncodeToMd5(bytes []byte) string {
	hashedBytes := md5.Sum(bytes)
	return fmt.Sprintf("%x", hashedBytes)
}
