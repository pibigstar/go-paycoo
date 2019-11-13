package paycoo

const (
	Format      = "JSON"
	Version     = "1.0"
	Charset     = "UTF-8"
	SignType    = "RSA"
	TimeFormat  = "2006-01-02 15:04:05"
	ContentType = "application/json"
)

const (
	DevAPI        = "http://open.wangtest.cn/api/gateway"
	ProductionAPI = "https://open.wangpos.com/api/gateway"
)

type PayParam interface {
	// API接口名
	Method() string
}
