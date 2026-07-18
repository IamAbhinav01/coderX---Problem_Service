package fomatter

import (
	"coderX/utils/json"
	"net/http"
)

func SucessResponse(w http.ResponseWriter,status int,message string,data any) error{

	response:=map[string]any{}
	response["status"] = status
	response["message"] = message
	response["data"] = data

	return json.ConvertTOJSON(w,status,response)
}