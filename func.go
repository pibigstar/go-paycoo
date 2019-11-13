package paycoo

// 推送订单至WPOS, https://www.yuque.com/paycoo/openapi/api-wpos.order.push2cashier
func (p *PayCoo) Push2cashier(param *Push2cashier) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}

// 唤醒WPOS扫描二维码, https://www.yuque.com/paycoo/openapi/api-wpos.cmd.qrscan
func (p *PayCoo) Qrscan(param *Qrscan) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}
