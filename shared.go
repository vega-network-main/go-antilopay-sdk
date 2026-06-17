package antilopay

import (
	"crypto/rsa"
	"net/http"
)

const (
	DefaultBaseURLv1 = "https://lk.antilopay.com/api/v1/"
	DefaultBaseURLv2 = "https://lk.antilopay.com/api/v2/"
)

type Client struct {
	projectID         string
	secretID          string
	privateKey        *rsa.PrivateKey
	callbackPublicKey *rsa.PublicKey
	httpClient        *http.Client

	baseURLv1 string
	baseURLv2 string
}

type ClientOption func(*Client)

func WithBaseURLv1(url string) ClientOption {
	return func(c *Client) {
		c.baseURLv1 = url
	}
}

func WithBaseURLv2(url string) ClientOption {
	return func(c *Client) {
		c.baseURLv2 = url
	}
}

type Customer struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	IP       string `json:"ip,omitempty"`
	FullName string `json:"fullname,omitempty"`
}

type PaymentParams struct {
	DirectNSPK bool `json:"direct_nspk,omitempty"`
}
