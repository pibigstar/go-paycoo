package paycoo

import (
	"testing"
)

func TestBarcodePay(t *testing.T) {
	param := &BarcodePay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "条码支付"
	param.TransAmount = "2.22"
	param.AuthCode = "2876344382566439"

	result, err := client.BarcodePay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestQRCodePay(t *testing.T) {
	param := &QRCodePay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "扫码支付下单"
	param.TransAmount = "2.22"
	param.PaymentMethod = "ALIPAY"

	result, err := client.QRCodePay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestH5Pay(t *testing.T) {
	param := &H5Pay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "H5支付下单"
	param.TransAmount = "2.22"
	param.ReturnURL = "https://pibigstar.github.io"

	result, err := client.H5Pay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestMiniPay(t *testing.T) {
	param := &MiniPay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "小程序支付"
	param.TransAmount = "2.22"
	param.Openid = "o-w4ZuO2yS033ZS18jan9TqghGTc"

	result, err := client.MiniPay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestAppPay(t *testing.T) {
	param := &AppPay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "APP支付"
	param.TransAmount = "2.22"
	param.PaymentMethod = "ALIPAY"

	result, err := client.AppPay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestWebPay(t *testing.T) {
	param := &WebPay{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "web支付"
	param.TransAmount = "2.22"
	param.PaymentMethod = "ALIPAY"

	result, err := client.WebPay(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestOrderQuery(t *testing.T) {
	param := &OrderQuery{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"

	result, err := client.OrderQuery(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestOrderRefund(t *testing.T) {
	param := &OrderRefund{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.OutRefundNo = "99999999"
	param.RefundDesc = "商品已售完"
	param.RefundAmount = "1.20"

	result, err := client.OrderRefund(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestPreAuthComp(t *testing.T) {
	param := &PreAuthComp{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.OutPeAuthCompNo = "999999"
	param.PreAuthCompAmount = "1.20"
	param.NotifyURL = "https://pibigstar.github.io"

	result, err := client.PreAuthComp(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
