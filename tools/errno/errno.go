package errno

import (
	"fmt"
	"github.com/wpf1118/toolbox/tools/logging"
)

type Error struct {
	Code    int64  `json:"errcode"`
	Message string `json:"message"`
}

func NewError(code int64, message string) Error {
	return Error{Code: code, Message: message}
}

func (e Error) AddF(v ...interface{}) Error {
	e.Message = fmt.Sprintf(e.Message, v...)

	return e
}

func (e Error) AddError(err error) Error {
	copyE := e
	copyE.Message = fmt.Sprintf("%s error: %v", e.Message, err)
	copyE.Log()

	return e
}

func (e Error) Log() Error {
	logging.ErrorF(e.Message)
	return e
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ErrorCode() int64 {
	return e.Code
}
