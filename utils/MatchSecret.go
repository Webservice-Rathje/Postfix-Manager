package utils

import (
	"encoding/base64"
	"io/ioutil"
)

func MatchSecret(secret string) bool {
	by, _ := ioutil.ReadFile("./SECRET_KEY")
	decoded, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return false
	}
	return equal(by, decoded)
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
