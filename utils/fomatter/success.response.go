package fomatter

import (
	"coderX/utils/json"
	"net/http"
)

func SucessResponse(w http.ResponseWriter,status int,message string,data any) error{

	response:=map[string]any{
		"Success": true,
		"Message": message,
	}
	if data != nil {
		response["data"] = data
	}
	if status == http.StatusCreated {
		response["error"] = map[string]any{}
	}

	return json.ConvertTOJSON(w,status,response)
}