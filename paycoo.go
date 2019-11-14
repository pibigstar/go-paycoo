package paycoo

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PayCoo struct {
	appId        string
	format       string // 仅支持JSON
	charset      string // 仅支持UTF-8
	signType     string // 仅支持RSA
	version      string // 固定值 1.0
	timestamp    string
	apiDomain    string
	isProduction bool
	privateKey   *rsa.PrivateKey
}

type Response struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Sign  string      `json:"sign"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
	Psn   string      `json:"psn"`
}

func NewClient(appId, privateKey string, isProduction bool) (*PayCoo, error) {
	client := &PayCoo{
		appId:        appId,
		format:       Format,
		charset:      Charset,
		signType:     SignType,
		version:      Version,
		isProduction: isProduction,
	}

	client.apiDomain = DevAPI
	if isProduction {
		client.apiDomain = ProductionAPI
	}

	key, err := parsePKCS8PrivateKey(formatPKCS8PrivateKey(privateKey))
	if err != nil {
		key, err = parsePKCS1PrivateKey(formatPKCS1PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}

	client.privateKey = key

	return client, nil
}

func (p *PayCoo) encodeParams(param PayParam) (url.Values, error) {
	values := url.Values{}
	values.Add("app_id", p.appId)
	values.Add("format", p.format)
	values.Add("charset", p.charset)
	values.Add("version", p.version)
	values.Add("method", param.Method())
	values.Add("timestamp", time.Now().Format(TimeFormat))

	m := make(map[string]interface{})
	bs, _ := json.Marshal(param)
	_ = json.Unmarshal(bs, &m)

	for key, value := range m {
		switch value.(type) {
		case string:
			values.Set(key, value.(string))
		default:
			bs, _ := json.Marshal(value)
			values.Set(key, string(bs))
		}
	}

	sign, err := sha256WithRSA(values, p.privateKey, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	values.Add("sign", sign)

	// sign_type 不参与签名运算
	values.Add("sign_type", p.signType)

	return values, nil
}

func (p *PayCoo) doRequest(param PayParam, result interface{}) error {
	var data io.Reader
	if param != nil {
		values, err := p.encodeParams(param)
		if err != nil {
			return err
		}
		var params = make(map[string]string, 0)
		for key := range values {
			value := strings.TrimSpace(values.Get(key))
			if len(value) > 0 {
				params[key] = values.Get(key)
			}
		}
		bs, _ := json.Marshal(params)
		fmt.Println(string(bs))
		data = bytes.NewReader(bs)
	}

	response, err := http.Post(p.apiDomain, ContentType, data)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, result)
	if err != nil {
		return err
	}
	return nil
}

func (p *PayCoo) AckNotification(w http.ResponseWriter) {
	success := struct {
		Code string `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: "0",
		Msg:  "SUCCESS",
	}
	data, _ := json.Marshal(success)
	w.Write(data)
}
