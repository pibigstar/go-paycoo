package paycoo

type PaperApply struct {
	SpId          string `json:"sp_id"`          // M 服务商ID
	MerchantName  string `json:"merchant_name"`  // M 商户名称
	PayChannel    string `json:"pay_channel"`    // M 支付通道, eg:JF621
	StoreNo       string `json:"store_no"`       // O 门店编号
	ThirdApplyNo  string `json:"third_apply_no"` // M 第三方申请流水号
	NotifyURL     string `json:"notify_url"`     // M 回调通知地址
	SalesmanName  string `json:"salesman_name"`  // M 销售人员姓名
	SalesmanPhone string `json:"salesman_phone"` // M 销售人员手机号
	Paper         Paper  `json:"paper"`          // M 进件详细信息
	FeeRates      string `json:"fee_rates"`      // 申请手续费率

}

type Paper struct {
	ApplyReason                string `json:"apply_reason"`                 // O 申请说明
	OpenExtraBiz               string `json:"open_extra_biz"`               // O 申请开通附加业务, eg: QR_PRE_AUTH
	MerchantAcquiringAgreement []byte `json:"merchant_acquiring_agreement"` // C 收单协议, JF621必传

	BusinessLicenseType          string `json:"business_license_type"`           // M 营业执照类型, 1: 三证合一
	BusinessLicenseCode          string `json:"business_license_code"`           // M 营业执照号
	BusinessLicenseName          string `json:"business_license_name"`           // M 商户经营名称
	BusinessLicensePhoto         []byte `json:"business_license_photo"`          // M 营业执照照片
	CategoryCode                 string `json:"category_code"`                   // M 行业类目
	BusinessContent              string `json:"business_content"`                // O 实体经营内容
	MerchantType                 string `json:"merchant_type"`                   // M 商户类型, 1: 个体工商户, 2: 企业, 3: 个人
	StoreName                    string `json:"store_name"`                      // M 店铺门面名称
	RegProvince                  string `json:"reg_province"`                    // M 省/自治区/直辖市
	RegCity                      string `json:"reg_city"`                        // M 市
	RegArea                      string `json:"reg_area"`                        // M 区
	RegAddressDetail             string `json:"reg_address_detail"`              // M 与经营地址一致
	BusinessLocationShopPhoto    []byte `json:"business_location_shop_photo"`    // M 门店外景照
	BusinessLocationHeadPhoto    []byte `json:"business_location_head_photo"`    // M 门店门头照
	BusinessLocationHallPhoto    []byte `json:"business_location_hall_photo"`    // M 门店内景照
	BusinessLocationCashierPhoto []byte `json:"business_location_cashier_photo"` // M 门店收银台照
	LegalPersonName              string `json:"legal_person_name"`               // M 法人姓名
	LegalPersonPhone             string `json:"legal_person_phone"`              // M 法人手机号
	LegalPersonCertType          string `json:"legal_person_cert_type"`          // M 法人证件类型, 1: 身份证
	LegalPersonCertId            string `json:"legal_person_cert_id"`            // M 法人证件ID
}
