package paycoo

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
)

const (
	PublicKeyPrefix = "-----BEGIN PUBLIC KEY-----"
	PublicKeySuffix = "-----END PUBLIC KEY-----"

	PKCS1Prefix = "-----BEGIN RSA PRIVATE KEY-----"
	PKCS1Suffix = "-----END RSA PRIVATE KEY-----"

	PKCS8Prefix = "-----BEGIN PRIVATE KEY-----"
	PKCS8Suffix = "-----END PRIVATE KEY-----"
)

var (
	ErrPrivateKey = errors.New("private key is empty")
)

func ParsePKCS1PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
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

func ParsePKCS8PrivateKey(data []byte) (key *rsa.PrivateKey, err error) {
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

func FormatPKCS1PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, PKCS8Prefix, "", 1)
	raw = strings.Replace(raw, PKCS8Suffix, "", 1)
	return formatKey(raw, PKCS1Prefix, PKCS1Suffix, 64)
}

func FormatPKCS8PrivateKey(raw string) []byte {
	raw = strings.Replace(raw, PKCS1Prefix, "", 1)
	raw = strings.Replace(raw, PKCS1Suffix, "", 1)
	return formatKey(raw, PKCS8Prefix, PKCS8Suffix, 64)
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
