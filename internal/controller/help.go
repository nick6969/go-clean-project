package controller

type GeneralSuccessResponse struct {
	Data interface{} `json:"data"`
}

type GeneralErrorResponse struct {
	Error string `json:"error"`
}
