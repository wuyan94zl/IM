package utils

type SuccessTmp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorTmp struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func Success(data interface{}) *SuccessTmp {
	return &SuccessTmp{
		Code: 200,
		Data: data,
	}
}

func Error(code int, message interface{}) *ErrorTmp {
	return &ErrorTmp{
		Code:    code,
		Message: message,
	}
}
