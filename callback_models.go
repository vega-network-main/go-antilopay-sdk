package main

type CallbackPayload struct {
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	CTime       string  `json:"ctime"`
	Currency    string  `json:"currency,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Fee         float64 `json:"fee,omitempty"`
	Description string  `json:"description,omitempty"`

	// Payment Specific Fields
	PaymentID         string    `json:"payment_id,omitempty"`
	OrderID           string    `json:"order_id,omitempty"`
	OriginalAmount    float64   `json:"original_amount,omitempty"`
	ProductName       string    `json:"product_name,omitempty"`
	PayMethod         string    `json:"pay_method,omitempty"`
	PayData           string    `json:"pay_data,omitempty"`
	CustomerIP        string    `json:"customer_ip,omitempty"`
	CustomerUserAgent string    `json:"customer_useragent,omitempty"`
	Customer          *Customer `json:"customer,omitempty"`
	MerchantExtra     string    `json:"merchant_extra,omitempty"`
	WithdrawID        string    `json:"withdraw_id,omitempty"`
	FeeType           string    `json:"fee_type,omitempty"`
	ProvideMethod     string    `json:"provide_method,omitempty"`
	RefundID          string    `json:"refund_id,omitempty"`
	TopupID           string    `json:"topup_id,omitempty"`
	CompleteTime      string    `json:"complete_time,omitempty"`
	AmountPaid        float64   `json:"amount_paid,omitempty"`
	TopupAmount       float64   `json:"topup_amount,omitempty"`
	SteamAccount      string    `json:"steam_account,omitempty"`
	Profit            float64   `json:"profit,omitempty"`
}
