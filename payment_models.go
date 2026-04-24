package antilopay

type PaymentCreateRequest struct {
	ProjectIdentificator string         `json:"project_identificator"`
	Amount               float64        `json:"amount"`
	OrderID              string         `json:"order_id"`
	Currency             string         `json:"currency"`
	ProductName          string         `json:"product_name"`
	ProductType          string         `json:"product_type"`
	ProductQuantity      int            `json:"product_quantity,omitempty"`
	VAT                  int            `json:"vat,omitempty"`
	Description          string         `json:"description"`
	SuccessURL           string         `json:"success_url,omitempty"`
	FailURL              string         `json:"fail_url,omitempty"`
	Customer             Customer       `json:"customer"`
	PreferMethods        []string       `json:"prefer_methods,omitempty"`
	MerchantExtra        string         `json:"merchant_extra,omitempty"`
	Params               *PaymentParams `json:"params,omitempty"`
}

type PaymentCreateResponse struct {
	Code          int    `json:"code"`
	PaymentID     string `json:"payment_id,omitempty"`
	PaymentURL    string `json:"payment_url,omitempty"`
	DirectNSPK    bool   `json:"direct_nspk,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Error         string `json:"error,omitempty"`
}

type PaymentCheckRequest struct {
	ProjectIdentificator string `json:"project_identificator"`
	OrderID              string `json:"order_id"`
}

type PaymentCheckResponse struct {
	Code              int       `json:"code"`
	PaymentID         string    `json:"payment_id,omitempty"`
	OrderID           string    `json:"order_id,omitempty"`
	PaymentURL        string    `json:"payment_url,omitempty"`
	CTime             string    `json:"ctime,omitempty"`
	Amount            float64   `json:"amount,omitempty"`
	OriginalAmount    float64   `json:"original_amount,omitempty"`
	Fee               float64   `json:"fee,omitempty"`
	Status            string    `json:"status,omitempty"`
	Currency          string    `json:"currency,omitempty"`
	ProductName       string    `json:"product_name,omitempty"`
	MerchantExtra     string    `json:"merchant_extra,omitempty"`
	Description       string    `json:"description,omitempty"`
	PayMethod         string    `json:"pay_method,omitempty"`
	PayData           string    `json:"pay_data,omitempty"`
	CustomerIP        string    `json:"customer_ip,omitempty"`
	CustomerUserAgent string    `json:"customer_useragent,omitempty"`
	Customer          *Customer `json:"customer,omitempty"`
	Error             string    `json:"error,omitempty"`
}
