// Package responses is bundle of function to get base response
package responses

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// BaseResponse is type of default response
type BaseResponse struct {
	Success bool `json:"success"`
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	Logger *zap.Logger
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

func (r BaseResponse) ErrorResponse(success bool, err error, message string) BaseResponse {
	if r.Logger != nil {
		r.Logger.Error("Response Error",
			zap.Error(err),
		)
	}

	return BaseResponse {
		Success: success,
		Code: 500,
		Data: fiber.Map{},
		Message: message,
	}
}