package auth

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateSalt() string {
	salt := make([]byte, 32)
	io.ReadFull(rand.Reader, salt)
	return base64.URLEncoding.EncodeToString(salt)
}
