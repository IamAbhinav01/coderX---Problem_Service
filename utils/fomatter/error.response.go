package fomatter

import (
	"coderX/utils/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter,status int,message string,err error) error {
	response := map[string]any{}
	errDetails := any(nil)
	if err != nil {
		errDetails = err.Error()
	}

	if status < 500 {
		response["message"] = message
		response["details"] = errDetails
		response["data"] = map[string]any{}
	} else {
		response["success"] = false
		response["message"] = message
		response["error"] = errDetails
		response["data"] = map[string]any{}
	}
	
	return json.ConvertTOJSON(w, status, response)
}