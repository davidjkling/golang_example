package main

import (
	"net/http"

	"github.com/pluralsight/webservices/controllers"
)

// main program starts web server on localhost:3000
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
