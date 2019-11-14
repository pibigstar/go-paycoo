package paycoo

import (
	"testing"
	"time"
)

func TestDownloadBill(t *testing.T) {
	param := &DownloadBill{}

	param.MerchantNo = "10030461"
	param.TransDate = time.Now().Format("2006-01-02")

	result, err := client.DownloadBill(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
