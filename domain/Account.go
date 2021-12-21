package domain

type Account struct {
	AccountId int     `json:"Id"`
	Name      string  `json:"Name"`
	Balance   float64 `json:"Balance"`
	Active    bool    `json:"Active"`
}
