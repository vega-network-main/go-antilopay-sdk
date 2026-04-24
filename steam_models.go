package antilopay

type SteamAccountCheckRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	SteamAccount         string `json:"steam_account"`
}

type SteamTopupCreateRequest struct {
	ProjectIdentificator string   `json:"project_identificator"`
	Amount               float64  `json:"amount"`
	TopupAmount          float64  `json:"topup_amount"`
	OrderID              string   `json:"order_id"`
	Currency             string   `json:"currency"`
	SteamAccount         string   `json:"steam_account"`
	Description          string   `json:"description"`
	SuccessURL           string   `json:"success_url,omitempty"`
	FailURL              string   `json:"fail_url,omitempty"`
	Customer             Customer `json:"customer"`
}

type SteamTopupCreateResponse struct {
	Code       int    `json:"code"`
	TopupID    string `json:"topup_id,omitempty"`
	PaymentURL string `json:"payment_url,omitempty"`
	Error      string `json:"error,omitempty"`
}

type SteamTopupCheckRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	OrderID              string `json:"order_id"`
}

type SteamTopupCheckResponse struct {
	Code              int       `json:"code"`
	TopupID           string    `json:"topup_id,omitempty"`
	OrderID           string    `json:"order_id,omitempty"`
	PaymentURL        string    `json:"payment_url,omitempty"`
	Status            string    `json:"status,omitempty"`
	CTime             string    `json:"ctime,omitempty"`
	CompleteTime      string    `json:"complete_time,omitempty"`
	AmountPaid        float64   `json:"amount_paid,omitempty"`
	TopupAmount       float64   `json:"topup_amount,omitempty"`
	Currency          string    `json:"currency,omitempty"`
	SteamAccount      string    `json:"steam_account,omitempty"`
	Fee               float64   `json:"fee,omitempty"`
	Profit            float64   `json:"profit,omitempty"`
	Description       string    `json:"description,omitempty"`
	PayMethod         string    `json:"pay_method,omitempty"`
	CustomerIP        string    `json:"customer_ip,omitempty"`
	CustomerUserAgent string    `json:"customer_useragent,omitempty"`
	Customer          *Customer `json:"customer,omitempty"`
	Error             string    `json:"error,omitempty"`
}
