package message

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var SuccessResponse = Response{
	StatusCode: 0,
	StatusMsg:  "Success!",
}
