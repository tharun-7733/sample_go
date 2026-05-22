package shared

type CoinBalanceParams struct {
	Username string `schema:"username"`
}

type CoinBalanceResponse struct {
	Balance int64 `json:"balance"`
	Code    int   `json:"code"`
}