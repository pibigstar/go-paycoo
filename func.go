package paycoo

// 推送订单至WPOS, https://www.yuque.com/paycoo/openapi/api-wpos.order.push2cashier
func (p *PayCoo) Push2cashier(param *Push2cashier) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}
