package paycoo

// 银联无感支付业务签约
type ContractSign struct {
	StoreNo       string `json:"store_no"`              // M 门店编号
	PaymentMethod string `json:"payment_method"`        // M 支付方式, SENSELESS_PAY: 银联二维码
	ApplyNo       string `json:"apply_no"`              // M 签约申请编号
	SpId          string `json:"sp_id,omitempty"`       // C 服务商ID
	MerchantNo    string `json:"merchant_no,omitempty"` // C 商户号
	ReturnUrl     string `json:"return_url,omitempty"`  // O 商家H5页面接受支付成功回调的页面地址
}

func (*ContractSign) Method() string {
	return "senselesspay.contract.sign"
}

type ContractQuery struct {
	StoreNo       string `json:"store_no"`              // M 门店编号
	PaymentMethod string `json:"payment_method"`        // M 支付方式, SENSELESS_PAY: 银联二维码
	SpId          string `json:"sp_id,omitempty"`       // C 服务商ID
	MerchantNo    string `json:"merchant_no,omitempty"` // C 商户号
	ApplyNo       string `json:"apply_no"`              // C 签约申请编号
	ContractNo    string `json:"contract_no"`           // C 签约编号, apply_no，contract_no至少一个必填
}

func (*ContractQuery) Method() string {
	return "senselesspay.contract.query"
}

type ContractTerminate struct {
	StoreNo       string `json:"store_no"`              // M 门店编号
	PaymentMethod string `json:"payment_method"`        // M 支付方式, SENSELESS_PAY: 银联二维码
	SpId          string `json:"sp_id,omitempty"`       // C 服务商ID
	MerchantNo    string `json:"merchant_no,omitempty"` // C 商户号
	ContractNo    string `json:"contract_no"`           // M 签约编号
}

func (*ContractTerminate) Method() string {
	return "senselesspay.contract.terminate"
}

// 发起扣款
type Withhold struct {
	StoreNo          string   `json:"store_no"`                    // M 门店编号
	OutOrderNo       string   `json:"out_order_no"`                // M 商户订单编号
	TransAmount      string   `json:"trans_amount"`                // M 订单金额
	Description      string   `json:"description"`                 // M 订单描述
	PaymentMethod    string   `json:"payment_method"`              // M 支付方式, SENSELESS_PAY: 银联二维码
	ContractNo       string   `json:"contract_no"`                 // M 签约号
	RiskInfo         RiskInfo `json:"risk_info"`                   // M 风控信息
	SpId             string   `json:"sp_id,omitempty"`             // C 服务商ID, 服务商模式时必须提供
	MerchantNo       string   `json:"merchant_no,omitempty"`       // C 商户号, 商户模式时必须提供
	TerminalNo       string   `json:"terminal_no,omitempty"`       // O 推送的目标终端编号, 设备EN
	TransCurrency    string   `json:"trans_currency,omitempty"`    // O 币种, 不填写则默认为RMB
	TransType        string   `json:"trans_type,omitempty"`        // O 交易类型, 1：消费
	NotifyURL        string   `json:"notify_url,omitempty"`        // O 通知地址
	Attach           string   `json:"attach,omitempty"`            // O 附加信息
	Longitude        string   `json:"longitude,omitempty"`         // O 经度
	Latitude         string   `json:"latitude,omitempty"`          // O 纬度
	EffectiveMinutes string   `json:"effective_minutes,omitempty"` // O 设置订单有效分钟数，超出有效时长不支付，订单将被关闭，不能再进行支付,默认为5分钟

}

type RiskInfo struct {
	DeviceID            string `json:"deviceID"`            // M 设备标识, 安卓:IMEI 苹果: IDFV
	DeviceType          string `json:"deviceType"`          // M 每类设备对应一个整数值，取值范围从 1 至 99。 1:手机，2:平板，3:手表，4:PC
	SourceIP            string `json:"sourceIP"`            // M 绑卡设备所在的公网IP，可用于定位所属地区，不是wifi连接时的局域网IP
	AccountIdHash       string `json:"accountIdHash"`       // M 用来标识用户在智能设备上登录账号 ID
	AccountRegisterTime string `json:"accountRegisterTime"` // M 用户在应用服务方注册时间，格式：yyyyMMddHHmmss
}

func (*Withhold) Method() string {
	return "senselesspay.withhold"
}
