package antilopay

import "net/http"

func (c *Client) CreateWithdraw(req *WithdrawCreateRequest) (*WithdrawCreateResponse, error) {
	var result WithdrawCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"withdraw/create", req, &result)
	return &result, err
}

func (c *Client) CheckWithdraw(req *WithdrawCheckRequest) (*WithdrawCheckResponse, error) {
	var result WithdrawCheckResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"withdraw/check", req, &result)
	return &result, err
}
