package paycoo

type WPOSParams struct {
	StoreNo        string `json:"store_no"`                    // M 门店编号
	TerminalNo     string `json:"terminal_no"`                 // M 推送的目标终端编号, 设备EN
	AcceptCashier  string `json:"accept_cashier"`              // M 指定处理的收银类APP, WIP、CSP、TFWIP
	TransType      string `json:"trans_type"`                  // M 交易类型
	OutOrderNo     string `json:"out_order_no"`                // M 商户订单编号
	OrderAmount    string `json:"order_amount"`                // M 支付金额
	SpId           string `json:"sp_id,omitempty"`             // C 服务商ID
	MerchantNo     string `json:"merchant_no,omitempty"`       // C 商户号
	OrigOutOrderNo string `json:"orig_out_order_no,omitempty"` // C 原商户订单号
	PaymentMethod  string `json:"payment_method,omitempty"`    // O 支付方式
	NotifyURL      string `json:"notify_url,omitempty"`        // O 通知地址
	Attach         string `json:"attach,omitempty"`            // O 附加信息
	PushMessage    string `json:"push_message,omitempty"`      // O 推送消息的内容
	VoiceMessage   string `json:"voice_message,omitempty"`     // O 语音播报内容

}

// 扩展参数
type ExtensionParameters struct {
	PaymentType      string `json:"payment_type,omitempty"`      // C 当payment_method不存在时生效, 扫码支付: QRCODE 刷卡支付: BANKCARD
	AcceptTimeOut    string `json:"accept_time_out,omitempty"`   // C 终端受理时间与请求发起时间间隔不得超过请求受理超时时间，否则交易关闭。单位：秒
	TransInstalments string `json:"trans_instalments,omitempty"` // M 花呗交易分期数，可选值:3,6,12
	TransInstalment  string `json:"trans_instalment,omitempty"`  // O 是否支持分期, Y:支持, N:不支持
}

// 推送订单至WPOS
type Push2cashier struct {
	WPOSParams
	BizNo                    string `json:"biz_no,omitempty"`                     // O 业务标识号
	IgnoreSupervisorPassword string `json:"ignore_supervisor_password,omitempty"` // O 忽略主管密码检验
	Description              string `json:"description,omitempty"`                // O 商户订单的描述信息
}

func (*Push2cashier) Method() string {
	return "wpos.order.push2cashier"
}

// WPOS扫描二维码
type QRScan struct {
	WPOSParams
	AnalysisMode   string `json:"analysis_mode"`             // M 数据解析模式 RAW_DATA：原始数据 （即不解析直接返回） WX_INVOICE：微信发票抬头
	AnalysisParams string `json:"analysis_params,omitempty"` // O 特定数据解析模式下，需要额外提供的参数
}

func (*QRScan) Method() string {
	return "wpos.cmd.qrscan"
}
