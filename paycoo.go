package paycoo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
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
	Sing         string
	version      string // 固定值 1.0
	timestamp    time.Time
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

func NewClient(appId, privateKey string, isProduction bool) *PayCoo {
	client := &PayCoo{
		appId:        appId,
		format:       Format,
		charset:      Charset,
		signType:     SignType,
		version:      Version,
		timestamp:    time.Now(),
		isProduction: isProduction,
		Client:       http.DefaultClient,
	}

	client.apiDomain = DevAPI
	if isProduction {
		client.apiDomain = ProductionAPI
	}
	var (
		key *rsa.PrivateKey
		err error
	)
	key, err = ParsePKCS1PrivateKey(FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		key, err = ParsePKCS8PrivateKey(FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			panic(err)
		}
	}

	client.privateKey = key

	return client
}

func (p *PayCoo) encodeParams(param PayParam) url.Values {
	values := url.Values{}
	values.Add("app_id", p.appId)
	values.Add("format", p.format)
	values.Add("charset", p.charset)
	values.Add("sign_type", p.signType)
	values.Add("version", p.version)
	values.Add("timestamp", p.timestamp.Format(TimeFormat))
	values.Add("method", param.Method())

	sign, err := signWithRSA(values, p.privateKey)
	if err != nil {
		panic(err)
	}
	values.Add("sign", sign)

	for key, value := range param.Params() {
		values.Add(key, value)
	}

	return values
}

func signWithRSA(values url.Values, privateKey *rsa.PrivateKey) (string, error) {
	src := values.Encode()
	var h = crypto.SHA256.New()
	h.Write([]byte(src))
	var hashed = h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	sign := base64.StdEncoding.EncodeToString(sig)
	return sign, nil
}

func (p *PayCoo) doRequest(params PayParam, result interface{}) error {
	var data io.Reader
	if params != nil {
		values := p.encodeParams(params)
		body, _ := url.PathUnescape(values.Encode())
		data = strings.NewReader(body)
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
