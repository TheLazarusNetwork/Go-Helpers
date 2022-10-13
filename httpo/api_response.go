package httpo

import (
	"github.com/gin-gonic/gin"
)

// ApiResponse defines struct used for http response
type ApiResponse[T any] struct {
	// Custom status code
	StatusCode int `json:"status,omitempty"`

	Error string `json:"error,omitempty"`

	// Message in case of success
	Message string `json:"message,omitempty"`
	Payload T      `json:"payload,omitempty"`
}

// Sends ApiResponse with gin context and standard statusCode
func (apiRes *ApiResponse[T]) Send(c *gin.Context, statusCode int) {
	c.JSON(statusCode, apiRes)
}

// Sends ApiResponse with gin context and with customStatusCode as standard one
func (apiRes *ApiResponse[T]) SendD(c *gin.Context) {
	c.JSON(apiRes.StatusCode, apiRes)
}

// NewSuccessResponse returns ApiResponse for success
func NewSuccessResponse[T any](customStatusCode int, message string, payload T) *ApiResponse[T] {
	return &ApiResponse[T]{
		StatusCode: customStatusCode,
		Message:    message,
		Payload:    payload,
	}
}

// NewSuccessResponse returns ApiResponse for error
func NewErrorResponse(customStatusCode int, errorStr string) *ApiResponse[any] {
	return &ApiResponse[any]{
		StatusCode: customStatusCode,
		Error:      errorStr,
	}
}
