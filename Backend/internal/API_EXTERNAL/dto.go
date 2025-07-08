package API_EXTERNAL

type StocksListResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}
type StockItem struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type PaginatedStocksResponse struct {
	Stocks     []Stock `json:"stocks"`
	Total      int64   `json:"total"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
	TotalPages int     `json:"total_pages"`
}

type StocksQueryParams struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Ticker   string `json:"ticker"`
}

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
