package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func WithTestMode(testURLv1, testURLv2 string) ClientOption {
	return func(c *Client) {
		c.baseURLv1 = testURLv1
		c.baseURLv2 = testURLv2
	}
}

func WithCustomHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(secretID, privateKeyBase64 string, opts ...ClientOption) (*Client, error) {
	pkBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, err
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(pkBytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	client := &Client{
		secretID:   secretID,
		privateKey: rsaKey,
		httpClient: &http.Client{},
		baseURLv1:  DefaultBaseURLv1,
		baseURLv2:  DefaultBaseURLv2,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func (c *Client) sign(body []byte) (string, error) {
	hash := sha256.New()
	hash.Write(body)
	d := hash.Sum(nil)

	s, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, d)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(s), nil
}

func (c *Client) doRequest(method, fullURL string, payload interface{}, result interface{}) (*http.Response, error) {
	var bodyReader io.Reader
	var reqBody []byte

	if payload != nil {
		rawJSON, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		compactBody := new(bytes.Buffer)
		if err = json.Compact(compactBody, rawJSON); err != nil {
			return nil, err
		}
		reqBody = compactBody.Bytes()
		bodyReader = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Apay-Secret-Id", c.secretID)

	if reqBody != nil {
		signature, err := c.sign(reqBody)
		if err != nil {
			return nil, err
		}
		req.Header.Set("X-Apay-Sign", signature)
		req.Header.Set("X-Apay-Sign-Version", "1")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if result != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}
		if len(bodyBytes) > 0 {
			if err := json.Unmarshal(bodyBytes, result); err != nil {
				return resp, err
			}
		}
	}

	return resp, nil
}

func (c *Client) CheckSignature(payload interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"signature/check", payload, &result)
	return result, err
}
