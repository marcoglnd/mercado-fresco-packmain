package schemes

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type JSONBadReqResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong parameter"`
	Data    interface{} `json:"data"`
}
