package paycoo

// 提交商户资料进件
type PaperApply struct {
	SpId                string              `json:"sp_id"`                          // M 服务商ID
	MerchantName        string              `json:"merchant_name"`                  // M 商户名称
	PayChannel          string              `json:"pay_channel"`                    // M 支付通道, eg:JF621
	StoreNo             string              `json:"store_no,omitempty"`             // O 门店编号
	ThirdApplyNo        string              `json:"third_apply_no"`                 // M 第三方申请流水号
	NotifyURL           string              `json:"notify_url"`                     // M 回调通知地址
	SalesmanName        string              `json:"salesman_name"`                  // M 销售人员姓名
	SalesmanPhone       string              `json:"salesman_phone"`                 // M 销售人员手机号
	Paper               Paper               `json:"paper"`                          // M 进件详细信息
	FeeRates            FeeRates            `json:"fee_rates"`                      // M 申请手续费率
	ExtensionParameters ExtensionParameters `json:"extension_parameters,omitempty"` // O 扩展参数

}

type Paper struct {
	ApplyReason                  string `json:"apply_reason,omitempty"`                 // O 申请说明
	OpenExtraBiz                 string `json:"open_extra_biz,omitempty"`               // O 申请开通附加业务, eg: QR_PRE_AUTH
	MerchantAcquiringAgreement   []byte `json:"merchant_acquiring_agreement,omitempty"` // C 收单协议, JF621必传
	BusinessLicenseType          string `json:"business_license_type"`                  // M 营业执照类型, 1: 三证合一
	BusinessLicenseCode          string `json:"business_license_code"`                  // M 营业执照号
	BusinessLicenseName          string `json:"business_license_name"`                  // M 商户经营名称
	BusinessLicensePhoto         []byte `json:"business_license_photo"`                 // M 营业执照照片
	CategoryCode                 string `json:"category_code"`                          // M 行业类目
	BusinessContent              string `json:"business_content,omitempty"`             // O 实体经营内容
	MerchantType                 string `json:"merchant_type"`                          // M 商户类型, 1: 个体工商户, 2: 企业, 3: 个人
	StoreName                    string `json:"store_name"`                             // M 店铺门面名称
	RegProvince                  string `json:"reg_province"`                           // M 省/自治区/直辖市
	RegCity                      string `json:"reg_city"`                               // M 市
	RegArea                      string `json:"reg_area"`                               // M 区
	RegAddressDetail             string `json:"reg_address_detail"`                     // M 与经营地址一致
	BusinessLocationShopPhoto    []byte `json:"business_location_shop_photo"`           // M 门店外景照
	BusinessLocationHeadPhoto    []byte `json:"business_location_head_photo"`           // M 门店门头照
	BusinessLocationHallPhoto    []byte `json:"business_location_hall_photo"`           // M 门店内景照
	BusinessLocationCashierPhoto []byte `json:"business_location_cashier_photo"`        // M 门店收银台照
	LegalPersonName              string `json:"legal_person_name"`                      // M 法人姓名
	LegalPersonPhone             string `json:"legal_person_phone"`                     // M 法人手机号
	LegalPersonCertType          string `json:"legal_person_cert_type"`                 // M 法人证件类型, 1: 身份证
	LegalPersonCertId            string `json:"legal_person_cert_id"`                   // M 法人证件ID
	LegalPersonCertPhotoFront    []byte `json:"legal_person_cert_photo_front"`          // M 法人证件正面照
	LegalPersonCertPhotoBack     []byte `json:"legal_person_cert_photo_back"`           // M 法人证件背面照
	LegalPersonCertExpired       string `json:"legal_person_cert_expired"`              // M 法人证件到期时间
	ContactPersonName            string `json:"contact_person_name"`                    // M 联系人姓名
	ContactPersonPhone           string `json:"contact_person_phone"`                   // M 联系人手机号
	ServiceTel                   string `json:"service_tel"`                            // M 客服电话
	Email                        string `json:"email"`                                  // M 客服邮件号
	SettleAccountType            string `json:"settle_account_type"`                    // M 结算账户类型, 1：对公结算2：对私结算-法人账户3：对私结算-被授权人账户
	BankAccountNo                string `json:"bank_account_no"`                        // M 结算账户号
	BankAccountName              string `json:"bank_account_name"`                      // M 结算账户名
	AccountOpeningLicense        []byte `json:"account_opening_license,omitempty"`      // C 开户许可证
	BankCardPhotoFront           []byte `json:"bank_card_photo_front,omitempty"`        // C 银行正面照
	BankCardPhotoBack            []byte `json:"bank_card_photo_back,omitempty"`         // C 对私结算时必填
	OpenBank                     string `json:"open_bank"`                              // M 开户银行名称
	OpenSubBank                  string `json:"open_sub_bank"`                          // M 开户支行名称
	OpenBankCode                 string `json:"open_bank_code"`                         // M 开户银联行号
	ReservePhone                 string `json:"reserve_phone"`                          // M 银行预留手机号
	MerchantAuthLicense          []byte `json:"merchant_auth_license,omitempty"`        // C 商户授权书
	CertId                       string `json:"cert_id,omitempty"`                      // C 被授权人证件号码
	CertExpired                  string `json:"cert_expired,omitempty"`                 // C 被授权人证件到期日
	CertPhotoFront               []byte `json:"cert_photo_front,omitempty"`             // C 被授权人证件正面照
	CertPhotoBack                []byte `json:"cert_photo_back,omitempty"`              // C 被授权人证件背面照
}

type FeeRates struct {
	PayMode         string `json:"pay_mode"`                    // M 支付方式 1：微信 2：支付宝 3：银联二维码
	FeeRateType     string `json:"fee_rate_type,omitempty"`     // C 手续费率类型, 当pay_mode=3时，必须提供
	FeeRateValue    string `json:"fee_rate_value"`              // M 手续费扣率，单位 %
	DebitRateValue  string `json:"debit_rate_value,omitempty"`  // C 借记卡手续费率, 当pay_mode=3时，必须提供
	DebitFeeCapping string `json:"debit_fee_capping,omitempty"` // C 借记卡手续费封顶, 当pay_mode=3时，必须提供
}

func (*PaperApply) Method() string {
	return "merchant.paper.apply"
}

// 商户资料补件
type PaperUpdate struct {
	WoId                string              `json:"wo_id"`                          // M 工单ID
	PayChannel          string              `json:"pay_channel"`                    // M 支付通道类型
	Paper               Paper               `json:"paper"`                          // M 商户进件资料
	FeeRates            FeeRates            `json:"fee_rates"`                      // M 手续费扣率
	ExtensionParameters ExtensionParameters `json:"extension_parameters,omitempty"` // O 扩展参数
}

func (*PaperUpdate) Method() string {
	return "merchant.paper.update"
}
