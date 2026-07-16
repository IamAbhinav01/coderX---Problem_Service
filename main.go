package main

import (
	"coderX/app"
	"fmt"
)

func main() {
	application := app.NewApplication()
	err := application.Run()
	if err != nil{
		fmt.Println("Error while initialising the server")
	}
}