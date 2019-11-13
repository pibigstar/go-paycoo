package paycoo

import "testing"

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

func TestQRScan(t *testing.T) {
	param := &QRScan{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.AcceptCashier = "WIP"
	param.TransType = "PURCHASE"
	param.OutOrderNo = "6666"
	param.OrderAmount = "1.20"
	param.AnalysisMode = "RAW_DATA"

	result, err := client.QRScan(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
