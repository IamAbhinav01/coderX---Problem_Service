package middleware

import (
	"coderX/models"
	"coderX/utils/fomatter"
	"coderX/utils/json"
	"context"
	"net/http"
)


type contextKey string 

var PayloadContextKey contextKey = "payload"

func ProblemCreation(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			var payload models.Problem
			
			err := json.DecodeFromJSON(r,&payload)
			if err != nil{
				fomatter.ErrorResponse(w,http.StatusBadRequest,"Error occured while decoding the json",err)
				return 
			}

			reqContext := r.Context()
			ctx:= context.WithValue(reqContext,PayloadContextKey,payload)
			

			next.ServeHTTP(w,r.WithContext(ctx))
		})
}