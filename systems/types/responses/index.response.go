// Package responses is bundle of function to get base response
package responses

// BaseResponse is type of default response
type BaseResponse struct {
	Success bool
	Code int
	Data interface{}
	Message string
}

// SuccessResponse is function for default success response
func (r BaseResponse) SuccessResponse(success bool, data interface{}) BaseResponse {
	return BaseResponse {
		Success: success,
		Code: 200,
		Data: data,
		Message: "Proceed Successfully",
	}
}