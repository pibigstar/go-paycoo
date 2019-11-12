package paycoo

func (p *PayCoo) Push2cashier(param *Push2cashier) (result *Response, err error) {
	err = p.doRequest(param, &result)
	return
}
