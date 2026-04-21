package main

import (
	"crypto/rsa"
	"net/http"
)

const (
	DefaultBaseURLv1 = "https://lk.antilopay.com/api/v1/"
	DefaultBaseURLv2 = "https://lk.antilopay.com/api/v2/"
)

type Client struct {
	secretID   string
	privateKey *rsa.PrivateKey
	httpClient *http.Client

	baseURLv1 string
	baseURLv2 string
}

type ClientOption func(*Client)

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
