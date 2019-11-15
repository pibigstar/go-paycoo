package paycoo

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"net/url"
	"sort"
	"strings"
)

const (
	kPublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	kPublicKeySuffix = "-----END PUBLIC KEY-----"

	kPKCS1Prefix = "-----BEGIN RSA PRIVATE KEY-----"
	KPKCS1Suffix = "-----END RSA PRIVATE KEY-----"

	kPKCS8Prefix = "-----BEGIN PRIVATE KEY-----"
	KPKCS8Suffix = "-----END PRIVATE KEY-----"
)

var (
	ErrPrivateKey = errors.New("private key have error")
	ErrPublicKey  = errors.New("public key have error")
)

func signParams(values url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	var params []string
	for key := range values {
		value := strings.TrimSpace(values.Get(key))
		if len(value) > 0 {
			params = append(params, key+"="+value)
		}
	}
	sort.Strings(params)
	var src = strings.Join(params, "&")
	return sha256WithRsaWithBase64([]byte(src), privateKey, hash)
}

func sha256WithRsaWithBase64(bs []byte, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	sig, err := sha256WithRsa(bs, privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func sha256WithRsa(src []byte, key *rsa.PrivateKey, hash crypto.Hash) ([]byte, error) {
	var h = hash.New()
	h.Write(src)
	var hashed = h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, key, hash, hashed)
}

func parsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPrivateKey
	}

	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, err
}

func parsePKCS8PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPrivateKey
	}

	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key, ok := rawKey.(*rsa.PrivateKey)
	if ok == false {
		return nil, ErrPrivateKey
	}

	return key, err
}

func formatPKCS1PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, kPKCS8Prefix, "", 1)
	raw = strings.Replace(raw, KPKCS8Suffix, "", 1)
	return formatKey(raw, kPKCS1Prefix, KPKCS1Suffix, 64)
}

func formatPKCS8PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, kPKCS1Prefix, "", 1)
	raw = strings.Replace(raw, KPKCS1Suffix, "", 1)
	return formatKey(raw, kPKCS8Prefix, KPKCS8Suffix, 64)
}

func parsePublicKey(data []byte) (key *rsa.PublicKey, err error) {
	var block *pem.Block
	block, _ = pem.Decode(data)
	if block == nil {
		return nil, ErrPublicKey
	}

	var pubInterface interface{}
	pubInterface, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, ErrPublicKey
	}

	return key, err
}

func formatPublicKey(key string) []byte {
	return formatKey(key, kPublicKeyPrefix, kPublicKeySuffix, 64)
}

func formatKey(raw, prefix, suffix string, lineCount int) []byte {
	if raw == "" {
		return nil
	}
	raw = strings.Replace(raw, prefix, "", 1)
	raw = strings.Replace(raw, suffix, "", 1)
	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)
	raw = strings.Replace(raw, "\t", "", -1)

	var sl = len(raw)
	var c = sl / lineCount
	if sl%lineCount > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix + "\n")
	for i := 0; i < c; i++ {
		var b = i * lineCount
		var e = b + lineCount
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}
