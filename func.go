package paycoo

// 推送订单至WPOS, https://www.yuque.com/paycoo/openapi/api-wpos.order.push2cashier
func (p *PayCoo) Push2cashier(param *Push2cashier) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 唤醒WPOS扫描二维码, https://www.yuque.com/paycoo/openapi/api-wpos.cmd.qrscan
func (p *PayCoo) QRScan(param *QRScan) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 条码(付款码)支付, https://www.yuque.com/paycoo/openapi/api-pay.barcodepay
func (p *PayCoo) BarcodePay(param *BarcodePay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 扫码支付下单, https://www.yuque.com/paycoo/openapi/api-pay.qrpay
func (p *PayCoo) QRCodePay(param *QRCodePay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 公众号/JSAPI/H5支付下单, https://www.yuque.com/paycoo/openapi/api-pay.h5pay
func (p *PayCoo) H5Pay(param *H5Pay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 小程序支付, https://www.yuque.com/paycoo/openapi/api-pay.miniapppay
func (p *PayCoo) MiniPay(param *MiniPay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// Deprecated, 开发中 App支付, https://www.yuque.com/paycoo/openapi/api-pay.nativeapppay
func (p *PayCoo) AppPay(param *AppPay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// Deprecated, 开发中  Web支付, https://www.yuque.com/paycoo/openapi/api-pay.webpay
func (p *PayCoo) WebPay(param *WebPay) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 交易查询, https://www.yuque.com/paycoo/openapi/api-pay.order.query
func (p *PayCoo) OrderQuery(param *OrderQuery) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 交易退款, https://www.yuque.com/paycoo/openapi/api-pay.order.refund
func (p *PayCoo) OrderRefund(param *OrderRefund) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 预授权完成, https://www.yuque.com/paycoo/openapi/api-pay.order.preauthcompletion
func (p *PayCoo) PreAuthComp(param *PreAuthComp) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 提交商户资料进件, https://www.yuque.com/paycoo/openapi/api-merchant.apply
func (p *PayCoo) PaperApply(param *PaperApply) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 商户资料补件, https://www.yuque.com/paycoo/openapi/api-merchant.modify
func (p *PayCoo) PaperUpdate(param *PaperUpdate) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 获取对账单, https://www.yuque.com/paycoo/openapi/api-accounting.downloadbill
func (p *PayCoo) DownloadBill(param *DownloadBill) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}
