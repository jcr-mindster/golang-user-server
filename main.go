package main

import (
	"fmt"
	"net/http"

	"example.com/controllers"
)

func main() {
	controllers.RegisterControllers()
	port := 3000
	portString := fmt.Sprintf(":%v", port)
	fmt.Println("Starting on port: ", port)
	http.ListenAndServe(portString, nil)
	fmt.Println("Stopped")
}
