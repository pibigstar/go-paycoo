package paycoo

import "testing"

func TestPaperApply(t *testing.T) {
	param := &PaperApply{}

	param.StoreNo = "166130"
	param.MerchantName = "派大星商戶"
	param.PayChannel = "JF621"
	param.ThirdApplyNo = "666666"
	param.NotifyURL = "https://pibigstar.github.io"
	param.SalesmanName = "派大星"
	param.SalesmanPhone = "110"
	param.Paper = Paper{
		ApplyReason: "测试",
	}
	param.FeeRates = FeeRates{
		PayMode: "1",
	}

	result, err := client.PaperApply(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}

func TestPaperUpdate(t *testing.T) {
	param := &PaperUpdate{}

	param.WoId = "66666666"
	param.PayChannel = "JF621"
	param.Paper = Paper{
		ApplyReason: "测试",
	}
	param.FeeRates = FeeRates{
		PayMode: "1",
	}

	result, err := client.PaperUpdate(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
