package paycoo

import (
	"bytes"
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
	publicKey    *rsa.PublicKey
}

type Response struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Sign  string      `json:"sign"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
	Psn   string      `json:"psn"`
}

func NewClient(appId, privateKey, publicKey string, isProduction bool) (*PayCoo, error) {
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

	key, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	client.privateKey = key

	pubKey, err := ParsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	client.publicKey = pubKey

	return client, nil
}

func (p *PayCoo) encodeParams(param PayParam) (url.Values, string, error) {
	values := url.Values{}
	values.Add("app_id", p.appId)
	values.Add("format", p.format)
	values.Add("charset", p.charset)
	values.Add("version", p.version)
	values.Add("sign_type", p.signType)
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

	sign, src, err := signParams(values, p.privateKey)
	if err != nil {
		return nil, src, err
	}
	values.Add("sign", sign)

	return values, src, nil
}

func (p *PayCoo) doRequest(param PayParam, result interface{}) error {
	var (
		data io.Reader
	)

	if param != nil {
		values, _, err := p.encodeParams(param)
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

	var resp Response
	err = json.Unmarshal(bs, &resp)
	if err != nil {
		return err
	}

	// TODO: 验签
	if resp.Sign != "" {
		values := url.Values{}
		values.Add("code", resp.Code)
		values.Add("msg", resp.Msg)
		values.Add("total", fmt.Sprintf("%d", resp.Total))
		values.Add("psn", resp.Psn)

		m := make([]map[string]string, 0)
		bs, _ := json.Marshal(resp.Data)
		_ = json.Unmarshal(bs, &m)
		dataStr, _ := json.Marshal(m)

		values.Add("data", string(dataStr))
		src := ParseValues(values)
		fmt.Println(src)
		err = VerifySignWithKey([]byte(src), resp.Sign, p.publicKey)
		if err != nil {
			fmt.Println(err.Error())
		}
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
