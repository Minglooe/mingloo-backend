package http_responses

// ResponseWrapper is a generic wrapper for API responses
type ResponseWrapper struct {
	Success bool        `extensions:"x-order=0" json:"success" example:"true"`
	Data    interface{} `extensions:"x-order=1" json:"data"`
} //@name Response
