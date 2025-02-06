// gin merupakan framework yang digunakan untuk keperluan untuk http routing
// cara menginstall gin melalui github go get -u github.com/gin-gonic/gin
// File entry point
package main

import (
	"gin-framework/routers"
)

func main() {
	// PORT
	port := ":8080"

	routers.StartServer().Run(port)
}

