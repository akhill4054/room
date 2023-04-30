package schemas

type ErrorResponse struct {
	Message string `json:"message" binding:"required"`
	Code    int    `json:"code,omitempty"`
	Detail  string `json:"detail,omitempty"`
}
