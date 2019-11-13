package paycoo

import (
	"encoding/json"
)

type PayMethod int

const (
	UNKNOWN           PayMethod = iota
	ALIPAY                      // 支付宝
	WECHATPAY                   // 微信
	UNIONPAY_QRCODE             // 银联二维码
	UNIONPAY_BANKCARD           // 银联卡
	OVERSEAS_BANKCARD           // 境外银行卡
)

func (p PayMethod) Name() string {
	switch p {
	case ALIPAY:
		return "ALIPAY"
	case WECHATPAY:
		return "WECHATPAY"
	case UNIONPAY_QRCODE:
		return "UNIONPAY_QRCODE"
	case UNIONPAY_BANKCARD:
		return "UNIONPAY_BANKCARD"
	case OVERSEAS_BANKCARD:
		return "OVERSEAS_BANKCARD"
	default:
		return ""
	}
}

type Push2cashier struct {
	StoreNo                  string `json:"store_no"`                             // M 门店编号
	TerminalNo               string `json:"terminal_no"`                          // M 推送的目标终端编号, 设备EN
	AcceptCashier            string `json:"accept_cashier"`                       // M 指定处理的收银类APP, WIP、CSP、TFWIP
	TransType                string `json:"trans_type"`                           // M 交易类型
	OutOrderNo               string `json:"out_order_no"`                         // M 商户订单编号
	OrderAmount              string `json:"order_amount"`                         // M 支付金额
	SpId                     string `json:"sp_id,omitempty"`                      // C 服务商ID
	MerchantNo               string `json:"merchant_no,omitempty"`                // C 商户号
	OrigOutOrderNo           string `json:"orig_out_order_no,omitempty"`          // C 原商户订单号
	PaymentMethod            string `json:"payment_method,omitempty"`             // O 支付方式
	NotifyURL                string `json:"notify_url,omitempty"`                 // O 通知地址
	Attach                   string `json:"attach,omitempty"`                     // O 附加信息
	BizNo                    string `json:"biz_no,omitempty"`                     // O 业务标识号
	PushMessage              string `json:"push_message,omitempty"`               // O 推送消息的内容
	VoiceMessage             string `json:"voice_message,omitempty"`              // O 语音播报内容
	ExtensionParameters      string `json:"extension_parameters,omitempty"`       // O 扩展参数
	IgnoreSupervisorPassword string `json:"ignore_supervisor_password,omitempty"` // O 忽略主管密码检验
	Description              string `json:"description,omitempty"`                // O 商户订单的描述信息
}

func (*Push2cashier) Method() string {
	return "wpos.order.push2cashier"
}

func (p *Push2cashier) Params() map[string]string {
	m := make(map[string]string)
	bytes, _ := json.Marshal(p)
	_ = json.Unmarshal(bytes, &m)
	return m
}
