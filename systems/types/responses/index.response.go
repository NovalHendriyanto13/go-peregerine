package responses

type BaseResponse struct {
	Success bool
	Code int
	Data interface{}
	Message string
}

func (r BaseResponse) SuccessResponse(success bool, data interface{}) BaseResponse {
	return BaseResponse {
		Success: success,
		Code: 200,
		Data: data,
		Message: "Proceed Successfully",
	}
}