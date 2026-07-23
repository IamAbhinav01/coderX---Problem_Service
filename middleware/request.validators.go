package middleware

import (
	dto "coderX/DTO"
	"coderX/utils/fomatter"
	"coderX/utils/json"
	"coderX/utils/validators"
	"context"
	"net/http"
)


type contextKey string 

var PayloadContextKey contextKey = "payload"

func ProblemCreation(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			var payload dto.GenerateProblem
			
			err := json.DecodeFromJSON(r,&payload)
			if err != nil{
				fomatter.ErrorResponse(w,http.StatusBadRequest,"Error occured while decoding the json",err)
				return 
			}
			validationErr := validators.Validate.Struct(payload)
			if validationErr != nil{
				fomatter.ErrorResponse(w,http.StatusBadRequest,"Invalid Bad Request Payload",validationErr)
				return 
			}
			
			reqContext := r.Context()
			ctx:= context.WithValue(reqContext,PayloadContextKey,payload)
			

			next.ServeHTTP(w,r.WithContext(ctx))
		})
}
