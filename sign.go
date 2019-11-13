package paycoo

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"github.com/smartwalle/crypto4go"
	"net/url"
	"sort"
	"strings"
)

func signWithPKCS1v15(values url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	var params []string
	for key := range values {
		value := strings.TrimSpace(values.Get(key))
		if len(value) > 0 {
			params = append(params, key+"="+value)
		}
	}
	sort.Strings(params)
	var src = strings.Join(params, "&")
	sig, err := crypto4go.RSASignWithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}
