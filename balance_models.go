package antilopay

type BalanceRequest struct {
	ProjectIdentificator string `json:"project_identificator,omitempty"`
}

type BalanceV1 struct {
	Available float64 `json:"available"`
	Blocked   float64 `json:"blocked"`
	Withdraw  float64 `json:"withdraw"`
}

type ProjectBalanceV1 struct {
	ProjectIdentificator string     `json:"project_identificator"`
	RUB                  BalanceV1  `json:"rub"`
	USD                  *BalanceV1 `json:"usd,omitempty"`
}

type BalanceResponseV1 struct {
	Code                 int                `json:"code"`
	ProjectIdentificator string             `json:"project_identificator,omitempty"`
	RUB                  *BalanceV1         `json:"rub,omitempty"`
	USD                  *BalanceV1         `json:"usd,omitempty"`
	Balances             []ProjectBalanceV1 `json:"balances,omitempty"`
	Error                string             `json:"error,omitempty"`
}

type BalanceV2 struct {
	Available string `json:"available"`
	Blocked   string `json:"blocked"`
	Withdraw  string `json:"withdraw"`
}

type ProjectBalanceV2 struct {
	ProjectIdentificator string     `json:"project_identificator"`
	RUB                  BalanceV2  `json:"rub"`
	USD                  *BalanceV2 `json:"usd,omitempty"`
}

type BalanceResponseV2 struct {
	Code                 int                `json:"code"`
	ProjectIdentificator string             `json:"project_identificator,omitempty"`
	RUB                  *BalanceV2         `json:"rub,omitempty"`
	USD                  *BalanceV2         `json:"usd,omitempty"`
	Balances             []ProjectBalanceV2 `json:"balances,omitempty"`
	Error                string             `json:"error,omitempty"`
}
