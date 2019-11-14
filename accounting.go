package paycoo

// 获取对账单
type DownloadBill struct {
	SpId       string `json:"sp_id,omitempty"`       // C 服务商ID, 服务商模式时必须提供
	MerchantNo string `json:"merchant_no,omitempty"` // C 商户号, 商户模式时必须提供
	TransDate  string `json:"trans_date"`            // M 商户所在时区的日期，格式：YYYY-MM-DD
}

func (*DownloadBill) Method() string {
	return "accounting.downloadbill"
}
