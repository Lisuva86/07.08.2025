package main

import (
	"zip_archive/api"
	"zip_archive/controller"
)

func main() {
	controller := controller.New()
	a := api.Init()
	v1 := a.Group("/api/v1")
	api.RegisterUserHandlers(v1, *controller)
	a.Run(":8080")
	//

}
