package main

import (
	"fmt"
	"github.com/lewandrade/go-webservice/controllers"
	"net/http"
)

const port int = 3000

var addr string = fmt.Sprintf("localhost:%v", port)

func main() {
	controllers.RegisterControllers()
	StartServer(addr)
}

func StartServer(addr string) {
	fmt.Printf("Starting server on port %v...", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
