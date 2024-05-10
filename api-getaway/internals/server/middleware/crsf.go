package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"io"

	log "github.com/sirupsen/logrus"
)

const (
	CSRFHeader = "X-CSRF-Token"
	csrfSalt = "KbWaoi5xtDC3GEfBa9ovQdzOzXsuVU9I"
)

func MakeToken(sid string) string {
	hash := sha256.New()
	_, err := io.WriteString(hash, csrfSalt+sid)
	if err != nil {
		log.Errorf("Make CSRF Token: %v", err)
	}
	token := base64.RawStdEncoding.EncodeToString(hash.Sum(nil))
	return token
}
func ValidateToken(token string,sid string)bool  {
	trueToken := MakeToken(sid)
	return token == trueToken
	
}