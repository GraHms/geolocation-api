package api

type ErrorResponse struct {
	Reason  string `json:"reason"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewAPIErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
func (c *ErrorResponse) SetResourceNotFound() {

	c.Message = "Not Found"
	c.Code = "ERR001"
	c.Reason = "Item was not Found"

}

func (c *ErrorResponse) SetGeoLocationNotFound() {
	c.Message = "Not Found"
	c.Code = "ERR001"
	c.Reason = "No Geolocation found for this IP"
}

func (c *ErrorResponse) SetInternalServerError() {

	c.Code = "ERR002"
	c.Reason = "Internal Server Error"
	c.Message = "Server encountered an unexpected condition"

}

func (c *ErrorResponse) SetMethodNotAllowed() {

	c.Code = "ERR004"
	c.Reason = "Method Not Allowed"
	c.Message = "Request method is not allowed"

}

func (c *ErrorResponse) Get() *ErrorResponse {
	return c
}
