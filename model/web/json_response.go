package web

type JsonResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type ValidationMessage struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}
