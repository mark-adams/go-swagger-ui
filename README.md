# Go Swagger UI

[![Build Status](https://travis-ci.org/mark-adams/go-swagger-ui.svg?branch=master)](https://travis-ci.org/mark-adams/go-swagger-ui)
[![Documentation](https://godoc.org/github.com/mark-adams/go-swagger-ui?status.svg)](https://godoc.org/github.com/mark-adams/go-swagger-ui)

This package provides a simple handler for rendering the Swagger UI in a Go application

# Example

```
package main

import (
	"net/http"

	"github.com/mark-adams/go-swagger-ui"
)

func main() {
	// Swagger UI using the latest JS / CSS from unpkg and loading the spec from /swagger.json
	http.Handle("/docs", swaggerui.Handler())

	// Swagger UI using a custom title, spec URL, and version of Swagger UI
	http.Handle("/docs-petstore", swaggerui.CustomHandler(
		"API Documentation", "http://petstore.swagger.io/v2/swagger.json", "latest",
	))

	http.ListenAndServe(":8080", nil)
}

```
