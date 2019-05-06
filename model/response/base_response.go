package response

type BaseResponse struct {
	Success bool `json:"success"`
	Code int `json:"code"`
	Msg string `json:"msg"`
}
