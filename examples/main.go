package main

import (
	"net/http"

	"github.com/mark-adams/go-swagger-ui"
)

func main() {
	http.Handle("/docs", swaggerui.CustomHandler(
		"API Documentation", "http://petstore.swagger.io/v2/swagger.json", "latest",
	))
	http.ListenAndServe(":8080", nil)
}
