package routes

import "github.com/gin-gonic/gin"

// APIError is the JSON representation of errors
type APIError struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
}

// APIReply is the JSON representation of a reply from the API
type APIReply struct {
	Data interface{} `json:"data"`
}

// EmptyReply is an empty reply
var EmptyReply = APIReply{}

// InternalError is the default Error
var InternalError = APIError{
	ErrorCode:   "internalError",
	Description: "Processing the request failed. This is an internal error.",
}

// InvalidParams signalizes that the JSON parameters are invalid
var InvalidParams = APIError{
	ErrorCode:   "invalidParameters",
	Description: "The parameters passed to the function were invalid",
}

// Router contains all API Endpoints
var Router *gin.Engine

// CreateRouter creates the router
func CreateRouter() {
	Router = gin.Default()
}
