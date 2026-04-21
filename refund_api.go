package main

import "net/http"

func (c *Client) CreateRefund(req *RefundCreateRequest) (*RefundCreateResponse, error) {
	var result RefundCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"refund/create", req, &result)
	return &result, err
}

func (c *Client) CreateReverse(req *ReverseCreateRequest) (*RefundCreateResponse, error) {
	var result RefundCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"reverse/create", req, &result)
	return &result, err
}

func (c *Client) CheckRefund(req *RefundCheckRequest) (*RefundCheckResponse, error) {
	var result RefundCheckResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"refund/check", req, &result)
	return &result, err
}
