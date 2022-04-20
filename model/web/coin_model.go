package web

type CoinCreateRequest struct {
	ID       int    `json:"id"`
	Coinname string `json:"coinname" validate:"required,min=1,max=100"`
}

type CoinCreateResponse struct {
	ID       int    `json:"id"`
	Coinname string `json:"coinname"`
}
