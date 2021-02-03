package main

import (
	"packages/configuration/routes"
)

func main() {
	routes.IntitailizeRoutes()
	routes.StartServer()
}
