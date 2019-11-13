package paycoo

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/smartwalle/crypto4go"
)

type PayCoo struct {
	appId        string
	format       string // 仅支持JSON
	charset      string // 仅支持UTF-8
	signType     string // 仅支持RSA
	version      string // 固定值 1.0
	timestamp    string
	Client       *http.Client
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
		Client:       http.DefaultClient,
	}

	client.apiDomain = DevAPI
	if isProduction {
		client.apiDomain = ProductionAPI
	}

	key, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		key, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(privateKey))
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
	values.Add("sign_type", p.signType)
	values.Add("version", p.version)
	values.Add("timestamp", time.Now().Format(TimeFormat))
	values.Add("method", param.Method())

	for key, value := range param.Params() {
		values.Add(key, value)
	}

	sign, err := signWithPKCS1v15(values, p.privateKey, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	values.Add("sign", sign)

	return values, nil
}

func (p *PayCoo) doRequest(params PayParam, result interface{}) error {
	var data io.Reader
	if params != nil {
		values, err := p.encodeParams(params)
		if err != nil {
			return err
		}
		data = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest("POST", p.apiDomain, data)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", ContentType)

	response, err := p.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		return err
	}
	return nil
}
