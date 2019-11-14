package paycoo

import (
	"testing"
)

func TestContractSign(t *testing.T) {
	param := &ContractSign{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.PaymentMethod = "SENSELESS_PAY"
	param.ApplyNo = "666666666"
	param.ReturnUrl = "https://pibigstar.github.io"

	result, err := client.ContractSign(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestContractQuery(t *testing.T) {
	param := &ContractQuery{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.PaymentMethod = "SENSELESS_PAY"
	param.ApplyNo = "666666666"

	result, err := client.ContractQuery(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestContractTerminate(t *testing.T) {
	param := &ContractTerminate{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.PaymentMethod = "SENSELESS_PAY"
	param.ContractNo = "666666666"

	result, err := client.ContractTerminate(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestWithhold(t *testing.T) {
	param := &Withhold{}
	param.MerchantNo = "10030461"
	param.StoreNo = "166130"
	param.TerminalNo = "888888888"
	param.OutOrderNo = "6666"
	param.Description = "发起扣款"
	param.TransAmount = "2.22"
	param.PaymentMethod = "SENSELESS_PAY"
	param.ContractNo = "9999999999"
	param.RiskInfo = RiskInfo{
		DeviceID:            "IMEI",
		DeviceType:          "4",
		SourceIP:            "192.168.0.1",
		AccountIdHash:       "22222",
		AccountRegisterTime: "20060102150405",
	}

	result, err := client.Withhold(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
