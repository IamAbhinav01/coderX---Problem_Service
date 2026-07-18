package fomatter

import "net/http"

func ErrorResponse(w http.ResponseWriter,status int,message string,err error){
	response := map[string]any{}
	response["status"] = status
	response["message"] = message

	if err != nil{
		response["error"] = err.Error()
	}else{
		response["error"] = nil
	}
}