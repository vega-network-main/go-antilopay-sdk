package antilopay

import "net/http"

func (c *Client) CheckSteamAccount(req *SteamAccountCheckRequest) (int, error) {
	resp, err := c.doRequest(http.MethodPost, c.baseURLv1+"steam/account/check", req, nil)
	if err != nil {
		if resp != nil {
			return resp.StatusCode, err
		}
		return 0, err
	}
	return resp.StatusCode, nil
}

func (c *Client) CreateSteamTopup(req *SteamTopupCreateRequest) (*SteamTopupCreateResponse, error) {
	var result SteamTopupCreateResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"steam/topup/create", req, &result)
	return &result, err
}

func (c *Client) CheckSteamTopup(req *SteamTopupCheckRequest) (*SteamTopupCheckResponse, error) {
	var result SteamTopupCheckResponse
	_, err := c.doRequest(http.MethodPost, c.baseURLv1+"steam/topup/check", req, &result)
	return &result, err
}
