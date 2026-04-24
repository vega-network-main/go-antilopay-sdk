package antilopay

import "net/http"

func (c *Client) GetBalanceV1(req *BalanceRequest) (*BalanceResponseV1, error) {
	var result BalanceResponseV1
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"project/balance", req, &result)
	return &result, err
}

func (c *Client) GetBalanceV2(req *BalanceRequest) (*BalanceResponseV2, error) {
	var result BalanceResponseV2
	_, err := c.doRequest(http.MethodPost, c.baseURLv2+"project/balance", req, &result)
	return &result, err
}
