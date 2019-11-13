# 旺POS支付

## Doc

https://www.yuque.com/paycoo/openapi

## Reminder
- `M` : 必须
- `C` : 二选一
- `O` : 可选

## Quick start
### Install
```bash
go get -u github.com/pibigstar/go-paycoo
```
### Demo
```go
func TestPush2cashier(t *testing.T) {
	param := &Push2cashier{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.AcceptCashier = "WIP"
	param.TransType = "PURCHASE"
	param.OutOrderNo = "6666"
	param.OrderAmount = "1.20"

	result, err := client.Push2cashier(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
```