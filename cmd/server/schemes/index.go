package schemes

type JSONSuccessResult struct {
	Data interface{} `json:“data”`
}

type JSONBadReqResult struct {
	Error interface{} `json:“error”`
}
