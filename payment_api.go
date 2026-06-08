package antilopay

import (
	"fmt"
	"net/http"
)

func (c *Client) CreatePayment(req *PaymentCreateRequest) (*PaymentCreateResponse, error) {
	if req.ProjectIdentificator == "" {
		req.ProjectIdentificator = c.projectID
	}

	var result PaymentCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"payment/create", req, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 || result.Error != "" {
		return &result, fmt.Errorf("antilopay api error (code %d): %s", result.Code, result.Error)
	}

	return &result, nil
}

func (c *Client) CheckPayment(req *PaymentCheckRequest) (*PaymentCheckResponse, error) {
	if req.ProjectIdentificator == "" {
		req.ProjectIdentificator = c.projectID
	}

	var result PaymentCheckResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"payment/check", req, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != 0 || result.Error != "" {
		return &result, fmt.Errorf("antilopay api error (code %d): %s", result.Code, result.Error)
	}

	return &result, nil
}
