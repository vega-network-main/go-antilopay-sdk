package antilopay

type RefundCreateRequest struct {
	ProjectIdentificator string  `json:"project_identificator"`
	TransactionID        string  `json:"transaction_id"`
	OrderID              string  `json:"order_id,omitempty"`
	Amount               float64 `json:"amount"`
}

type RefundCreateResponse struct {
	Code     int    `json:"code"`
	RefundID string `json:"refund_id,omitempty"`
	Error    string `json:"error,omitempty"`
}

type RefundCheckRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	RefundID             string `json:"refund_id,omitempty"`
	OrderID              string `json:"order_id,omitempty"`
}

type RefundCheckResponse struct {
	Code      int     `json:"code"`
	RefundID  string  `json:"refund_id,omitempty"`
	OrderID   string  `json:"order_id,omitempty"`
	PaymentID string  `json:"payment_id,omitempty"`
	Status    string  `json:"status,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
	Error     string  `json:"error,omitempty"`
}

type ReverseCreateRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	TransactionID        string `json:"transaction_id"`
	OrderID              string `json:"order_id,omitempty"`
}
