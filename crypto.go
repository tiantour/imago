package imago

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

//Base64Encode base64加密
func (c *icrypto) Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//Base64Decode base64解密
func (c *icrypto) Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

//Md532 md5 32
func (c *icrypto) Md532(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//Md516 md5 16
func (c *icrypto) Md516(data string) string {
	return c.Md532(data)[8:24]
}

//SHA1 sha1
func (c *icrypto) SHA1(data string) []byte {
	h := sha1.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

//SHA256 sha256
func (c *icrypto) SHA256(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

// HmacSha1 HmacSha1
func (c *icrypto) HmacSha1(publicKey, privateKey string) []byte {
	mac := hmac.New(sha1.New, []byte(privateKey))
	mac.Write([]byte(publicKey))
	return mac.Sum(nil)
}

//Pbkdf2Sha256 Pbkdf2Sha256
func (c *icrypto) Pbkdf2Sha256(data, salt string) string {
	iterations := 12000
	dk := pbkdf2.Key([]byte(data), []byte(salt), iterations, 32, sha256.New)
	return fmt.Sprintf("pbkdf2_sha256$%s$%s$%s", "12000", salt, base64.StdEncoding.EncodeToString(dk))
}

//Encrypt rsa加密
func (c *icrypto) RSAEncrypt(origdata string) (result string, err error) {
	publicKey, err := File.Read("public/conf/alipay/rsa_public_key.pem")
	if err != nil {
		return
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	resultByte, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origdata))
	if err != nil {
		return
	}
	result = c.Base64Encode(resultByte)
	return
}

//Decrypt rsa解密
func (c *icrypto) RSADecrypt(ciphertext string) (result string, err error) {
	privateKey, err := File.Read("public/conf/future/rsa_private_key.pem")
	if err != nil {
		return
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error")
		return
	}
	privInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	resultTemp, err := c.Base64Decode(ciphertext)
	if err != nil {
		return
	}
	resultByte, err := rsa.DecryptPKCS1v15(rand.Reader, privInterface, resultTemp)
	if err != nil {
		return
	}
	result = string(resultByte)
	return
}

//Sign rsa签名
func (c *icrypto) RSASign(origdata string) (result string, err error) {
	privateKey, err := File.Read("public/conf/future/rsa_private_key.pem")
	if err != nil {
		return
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		err = errors.New("private key error")
		return
	}
	privInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	digest := c.SHA1(origdata)
	resultByte, err := rsa.SignPKCS1v15(rand.Reader, privInterface, crypto.SHA1, digest)
	if err != nil {
		return
	}
	result = c.Base64Encode(resultByte)
	return
}

//Verify rsa签名验证
func (c *icrypto) RSAVerify(origdata, ciphertext string) (status bool, err error) {
	publicKey, err := File.Read("public/conf/alipay/rsa_public_key.pem")
	if err != nil {
		return
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		err = errors.New("public key error")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	digest := c.SHA1(origdata)
	resultTemp, err := c.Base64Decode(ciphertext)
	if err != nil {
		return
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, resultTemp)
	if err != nil {
		return
	}
	status = true
	return
}
