package errno

import "fmt"

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func NewError(code int64, message string) Error {
	return Error{Code: code, Message: message}
}

func (e Error) AddF(v ...interface{}) Error {
	e.Message = fmt.Sprintf(e.Message, v...)

	return e
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ErrorCode() int64 {
	return e.Code
}
