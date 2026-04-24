package antilopay

import "net/http"

func (c *Client) CreatePayment(req *PaymentCreateRequest) (*PaymentCreateResponse, error) {
	var result PaymentCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"payment/create", req, &result)
	return &result, err
}

func (c *Client) CheckPayment(req *PaymentCheckRequest) (*PaymentCheckResponse, error) {
	var result PaymentCheckResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"payment/check", req, &result)
	return &result, err
}
