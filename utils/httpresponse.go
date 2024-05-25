package utils

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResponseBody struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func RenderJSONWithData(w http.ResponseWriter, r *http.Request, code int, data interface{}, err interface{}, message interface{}) {
	render.Status(r, code)
	render.JSON(w, r, ResponseBody{
		Status:  http.StatusText(code),
		Code:    code,
		Data:    data,
		Error:   err,
		Message: message,
	})
}
