package paycoo

type PayParams struct {
	StoreNo             string              `json:"store_no"`                       // M 门店编号
	TerminalNo          string              `json:"terminal_no"`                    // M 推送的目标终端编号, 设备EN
	OutOrderNo          string              `json:"out_order_no"`                   // M 商户订单编号
	TransAmount         string              `json:"trans_amount"`                   // M 订单金额
	Description         string              `json:"description"`                    // M 订单描述
	SpId                string              `json:"sp_id,omitempty"`                // C 服务商ID, 服务商模式时必须提供
	MerchantNo          string              `json:"merchant_no,omitempty"`          // C 商户号, 商户模式时必须提供
	GoodsDetail         string              `json:"goods_detail,omitempty"`         // O 订单包含的商品列表信息，json格式
	DiscountableAmount  string              `json:"discountable_amount,omitempty"`  // O 参与优惠计算的金额，取值范围[0.01,100000000]，支付宝可传递该参数
	TransCurrency       string              `json:"trans_currency,omitempty"`       // O 币种, 不填写则默认为RMB
	NotifyURL           string              `json:"notify_url,omitempty"`           // O 通知地址
	Attach              string              `json:"attach,omitempty"`               // O 附加信息
	Longitude           string              `json:"longitude,omitempty"`            // O 经度
	Latitude            string              `json:"latitude,omitempty"`             // O 纬度
	EffectiveMinutes    string              `json:"effective_minutes,omitempty"`    // O 设置订单有效分钟数，超出有效时长不支付，订单将被关闭，不能再进行支付,默认为5分钟
	ExtensionParameters ExtensionParameters `json:"extension_parameters,omitempty"` // O 扩展输入参数，后续定义增加的参数存储于此JSON可变结构中
}

// 条码(付款码)支付
type BarcodePay struct {
	PayParams
	AuthCode  string `json:"auth_code"`            // M 扫码支付授权码 示例：支付宝- 2876344382566439
	TransType string `json:"trans_type,omitempty"` // O 交易类型 1: 消费, 2: 预授权
}

func (*BarcodePay) Method() string {
	return "pay.barcodepay"
}

// 扫码支付下单
type QRCodePay struct {
	PayParams
	PaymentMethod string `json:"payment_method"`       // M 支付方式
	TransType     string `json:"trans_type,omitempty"` // O 交易类型 1: 消费, 2: 预授权
}

func (*QRCodePay) Method() string {
	return "pay.qrcodepay"
}

// 公众号/JSAPI/H5支付下单
type H5Pay struct {
	PayParams
	ReturnURL string `json:"return_url,omitempty"` // O 前台回调地址
}

func (*H5Pay) Method() string {
	return "pay.h5pay"
}

// 小程序支付
type MiniPay struct {
	PayParams
	Openid string `json:"openid"` // M 微信用户ID
}

func (*MiniPay) Method() string {
	return "pay.minipay"
}

// APP支付
type AppPay struct {
	PayParams
	PaymentMethod string `json:"payment_method"` // M 支付方式
}

func (*AppPay) Method() string {
	return "pay.apppay"
}

// Web网页支付
type WebPay struct {
	PayParams
	PaymentMethod string `json:"payment_method"`         // M 支付方式
	BrowserType   string `json:"browser_type,omitempty"` // O 浏览器类型, PC、WAP
}

func (*WebPay) Method() string {
	return "pay.webpay"
}

// 交易查询
type OrderQuery struct {
	StoreNo    string `json:"store_no"`                  // M 门店编号
	TerminalNo string `json:"terminal_no"`               // M 推送的目标终端编号, 设备EN
	OutOrderNo string `json:"out_order_no,omitempty"`    // C 商户订单编号
	TransNo    string `json:"trans_no,omitempty"`        // C 交易号
	SpId       string `json:"sp_id,omitempty,omitempty"` // C 服务商ID, 服务商模式时必须提供
	MerchantNo string `json:"merchant_no,omitempty"`     // C 商户号, 商户模式时必须提供
}

func (*OrderQuery) Method() string {
	return "pay.orderquery"
}

// 交易退款
type OrderRefund struct {
	StoreNo      string `json:"store_no"`               // M 门店编号
	TerminalNo   string `json:"terminal_no"`            // M 推送的目标终端编号, 设备EN
	RefundAmount string `json:"refund_amount"`          // M 退款金额
	OutRefundNo  string `json:"out_refund_no"`          // M 退款单号, 同一退款单号多次请求只退一笔， 同一app_id下商户退款单号不能重复
	OutOrderNo   string `json:"out_order_no,omitempty"` // C 商户订单编号
	TransNo      string `json:"trans_no,omitempty"`     // C 交易号
	SpId         string `json:"sp_id,omitempty"`        // C 服务商ID, 服务商模式时必须提供
	MerchantNo   string `json:"merchant_no,omitempty"`  // C 商户号, 商户模式时必须提供
	RefundDesc   string `json:"refund_desc,omitempty"`  // O 退款原因
}

func (*OrderRefund) Method() string {
	return "pay.orderrefund"
}

type PreAuthComp struct {
	StoreNo           string `json:"store_no"`              // M 门店编号
	TerminalNo        string `json:"terminal_no"`           // M 推送的目标终端编号, 设备EN
	OutOrderNo        string `json:"out_order_no"`          // M 商户订单编号
	OutPeAuthCompNo   string `json:"out_preauthcomp_no"`    // M 商户预授权完成订单号
	PreAuthCompAmount string `json:"preauthcomp_amount"`    // M 本次预授权完成金额
	SpId              string `json:"sp_id,omitempty"`       // C 服务商ID, 服务商模式时必须提供
	MerchantNo        string `json:"merchant_no,omitempty"` // C 商户号, 商户模式时必须提供
	NotifyURL         string `json:"notify_url,omitempty"`  // O 后台通知地址
}

func (*PreAuthComp) Method() string {
	return "pay.preauth.comp"
}
