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

// crypto crypto
type crypto struct{}

// Crypto new crypto
var Crypto = new(crypto)

//Base64Encode crypto base64 encode
// date 2016-12-31
// author andy.jiang
func (c crypto) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//Base64Decode crypto base64 decode
// date 2016-12-31
// author andy.jiang
func (c crypto) Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

//Md532 crypto md532
// date 2016-12-31
// author andy.jiang
func (c crypto) Md532(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//Md516 crypto 16
// date 2016-12-31
// author andy.jiang
func (c crypto) Md516(data string) string {
	return c.Md532(data)[8:24]
}

//SHA1 crypto sha1
// date 2016-12-31
// author andy.jiang
func (c crypto) SHA1(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

//SHA256 crypto sha256
// date 2016-12-31
// author andy.jiang
func (c crypto) SHA256(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

// HmacSha1 crypto hmacsha1
// date 2016-12-31
// author andy.jiang
func (c crypto) HmacSha1(publicKey, privateKey string) []byte {
	mac := hmac.New(sha1.New, []byte(privateKey))
	mac.Write([]byte(publicKey))
	return mac.Sum(nil)
}

// Pbkdf2Sha256 crypto pbkdf2sha256
// date 2016-12-31
// author andy.jiang
func (c crypto) Pbkdf2Sha256(data, salt string, iterations int) string {
	dk := pbkdf2.Key([]byte(data), []byte(salt), iterations, 32, sha256.New)
	return fmt.Sprintf("pbkdf2_sha256$%d$%s$%s", iterations, salt, base64.StdEncoding.EncodeToString(dk))
}
