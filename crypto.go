package imago

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

//Base64Encode
func (c *iCrypto) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//Base64Decode
func (c *iCrypto) Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

//Md532
func (c *iCrypto) Md532(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//Md516
func (c *iCrypto) Md516(data string) string {
	return c.Md532(data)[8:24]
}

//SHA1
func (c *iCrypto) SHA1(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

//SHA256
func (c *iCrypto) SHA256(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

// HmacSha1
func (c *iCrypto) HmacSha1(publicKey, privateKey string) []byte {
	mac := hmac.New(sha1.New, []byte(privateKey))
	mac.Write([]byte(publicKey))
	return mac.Sum(nil)
}

// Pbkdf2Sha256
func (c *iCrypto) Pbkdf2Sha256(data, salt string) string {
	iterations := 12000
	dk := pbkdf2.Key([]byte(data), []byte(salt), iterations, 32, sha256.New)
	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", iterations, salt, base64.StdEncoding.EncodeToString(dk))
}
