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
	w.WriteHeader(http.StatusOK)
}

type Notification struct {
	AppId              string `json:"app_id"`
	Format             string `json:"format"`
	Charset            string `json:"charset"`
	SignType           string `json:"sign_type"`
	Sign               string `json:"sign"`
	Version            string `json:"version"`
	Timestamp          string `json:"timestamp"`
	SpId               string `json:"sp_id,omitempty"`                 // C 服务商ID
	MerchantNo         string `json:"merchant_no,omitempty"`           // C 商户号
	StoreNo            string `json:"store_no,omitempty"`              // M 门店编号
	TerminalNo         string `json:"terminal_no,omitempty"`           // M 推送的目标终端编号, 设备EN
	TransNo            string `json:"trans_no,omitempty"`              // M 交易号
	OutOrderNo         string `json:"out_order_no,omitempty"`          // M 商户订单编号
	PayPlatformTransNo string `json:"pay_platform_trans_no,omitempty"` // 支付平台交易号
	TransStatus        string `json:"trans_status,omitempty"`          // M 交易类型
	PaymentMethod      string `json:"payment_method,omitempty"`        // O 支付方式
	TransType          string `json:"trans_type,omitempty"`            // M 交易类型
	PayUserAccountId   string `json:"pay_user_account_id,omitempty"`   // M 卖家账号标识
	TransCurrency      string `json:"trans_currency,omitempty"`        // O 标价币种
	ExchangeRate       string `json:"exchange_rate,omitempty"`         // O 汇率
	TransAmount        string `json:"trans_amount,omitempty"`          // M 交易金额
	CustomerPaidAmount string `json:"customer_paid_amount,omitempty"`  // M 顾客实付金额
	DiscountBmopc      string `json:"discount_bmopc,omitempty"`        // O 支付通道商户优惠金额
	DiscountBpc        string `json:"discount_bpc,omitempty"`          // O 支付通道优惠金额
	TransEndTime       string `json:"trans_end_time,omitempty"`        // M 交易完成时间
	Attach             string `json:"attach,omitempty,omitempty"`      // O 附加信息
	CardAttr           string `json:"card_attr,omitempty"`             // O 卡属性
}

func (p *PayCoo) GetNotification(req *http.Request) (*Notification, error) {
	result := &Notification{}
	err := json.NewDecoder(req.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	bs, _ := json.Marshal(result)
	var resultMap map[string]string
	_ = json.Unmarshal(bs, &resultMap)

	str := buildSignStr(resultMap)
	// 验签
	err = VerifySignWithKey([]byte(str), result.Sign, p.publicKey)
	if err != nil {
		return result, err
	}
	return result, nil
}

func buildSignStr(resultMap map[string]string) string {
	urls := url.Values{}
	for key, value := range resultMap {
		if value == "" {
			continue
		}
		urls.Add(key, value)
	}
	return ParseValues(urls)
}
