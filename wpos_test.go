package paycoo

import (
	"fmt"
	"testing"
	"time"
)

func TestPush2cashier(t *testing.T) {
	param := &Push2cashier{}
	param.StoreNo = "166130"
	param.TerminalNo = "8888888"
	param.AcceptCashier = "CSP"
	param.TransType = "PURCHASE"
	param.OutOrderNo = fmt.Sprintf("%d", time.Now().Unix())
	param.OrderAmount = "0.01"

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
	param.TerminalNo = "8888888"
	param.AcceptCashier = "CSP"
	param.TransType = "PURCHASE"
	param.OutOrderNo = fmt.Sprintf("%d", time.Now().Unix())
	param.OrderAmount = "0.01"
	param.AnalysisMode = "RAW_DATA"

	result, err := client.QRScan(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
