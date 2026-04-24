package antilopay

type WithdrawCreateRequest struct {
	ProjectIdentificator string  `json:"project_identificator"`
	OrderID              string  `json:"order_id,omitempty"`
	Amount               float64 `json:"amount"`
	Method               string  `json:"method"`
	Account              string  `json:"account"`
	FeeType              string  `json:"fee_type,omitempty"`
}

type WithdrawCreateResponse struct {
	Code       int    `json:"code"`
	WithdrawID string `json:"withdraw_id,omitempty"`
	Error      string `json:"error,omitempty"`
}

type WithdrawCheckRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	WithdrawID           string `json:"withdraw_id,omitempty"`
	OrderID              string `json:"order_id,omitempty"`
}

type WithdrawCheckResponse struct {
	Code          int     `json:"code"`
	WithdrawID    string  `json:"withdraw_id,omitempty"`
	OrderID       string  `json:"order_id,omitempty"`
	CTime         string  `json:"ctime,omitempty"`
	Status        string  `json:"status,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Fee           float64 `json:"fee,omitempty"`
	FeeType       string  `json:"fee_type,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	ProvideMethod string  `json:"provide_method,omitempty"`
	Error         string  `json:"error,omitempty"`
}
