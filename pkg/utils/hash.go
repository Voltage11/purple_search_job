package utils

import (
	"crypto/sha256"
	"fmt"
)

func StrToHash(strForHash string) string {
	hasher := sha256.New()
    hasher.Write([]byte(strForHash))
    return fmt.Sprintf("%x", hasher.Sum(nil)) 
}