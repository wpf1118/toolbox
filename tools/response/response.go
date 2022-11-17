package response

import (
	"encoding/json"
	"net/http"

	"github.com/wpf1118/toolbox/tools/errno"
)

// Respond setups the response correctly for HTTP requests
func Respond(w http.ResponseWriter, code int, payload interface{}) {
	if payload == nil {
		payload = NewResSuccess("success")
	}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, err errno.Error, payload ...interface{}) {
	Respond(w, http.StatusOK, NewResError(err, payload...))
}

func Ok(w http.ResponseWriter, payload interface{}) {
	Respond(w, http.StatusOK, NewResSuccess(payload))
}

type ResData struct {
	Success      bool        `json:"success"`
	ErrorCode    int64       `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

type List struct {
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	List  interface{} `json:"list"`
}

func NewResSuccess(data interface{}) *ResData {
	return &ResData{
		Success:      true,
		Data:         data,
		ErrorCode:    0,
		ErrorMessage: "",
	}
}

func NewResError(err errno.Error, args ...interface{}) *ResData {
	var data interface{}
	if len(args) >= 1 {
		data = args[0]
	}

	return &ResData{
		Success:      true,
		ErrorCode:    err.ErrorCode(),
		ErrorMessage: err.Error(),
		Data:         data,
	}
}
