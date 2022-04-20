package web

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ValidationMessage struct {
	FailedField string
	Tag         string
	Value       string
}
