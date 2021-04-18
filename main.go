package main

import (
	"github.com/lewandrade/go-webservice/controllers"
	"net/http"
)

func main() {
	//u := models.User{
	//	ID:        2,
	//	FirstName: "Lew",
	//	LastName:  "Andrade",
	//}
	//
	//fmt.Println(u)
	//port := 3000
	//_, err := startWebServer(port, 5)
	//fmt.Println(port, err)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
