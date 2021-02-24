package utils

import (
	"crypto/rand"
	"os"
)

func GenerateSecret() {
	if _, err := os.Stat("./SECRET_KEY"); os.IsNotExist(err) {
		os.Create("./SECRET_KEY")
		secret := make([]byte, 4096)
		rand.Read(secret)
		f, _ := os.OpenFile("./SECRET_KEY", os.O_WRONLY, 0644)
		f.Write(secret)
	}
}
