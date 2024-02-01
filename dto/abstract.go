package dto

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
