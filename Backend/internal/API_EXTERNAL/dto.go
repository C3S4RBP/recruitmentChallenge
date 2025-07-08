package API_EXTERNAL

type StockResponse struct {
	Ticker     string `gorm:"primaryKey" json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StocksListResponse struct {
	Items    []StockResponse `json:"items"`
	NextPage string          `json:"next_page"`
}

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
