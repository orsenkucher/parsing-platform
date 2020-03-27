package hsh

import (
	"crypto/sha256"
	"encoding/base64"
)

var Hashs map[string]string = make(map[string]string)

func Hash(key string) string {
	bytes := []byte(key)
	sha := sha256.New()
	_, err := sha.Write(bytes)
	if err != nil {
		panic(err)
	}
	bytes = sha.Sum(nil)
	res := base64.StdEncoding.EncodeToString(bytes)
	Hashs[res] = key
	return res
}
