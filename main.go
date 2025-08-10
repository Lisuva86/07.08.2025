package main

import (
	"zip_archive/api"
	"zip_archive/controller"
	"zip_archive/entity"
)

func main() {
	controller := controller.New()
	//создать папок для скачанных файлов и архивов
	controller.CreateFolder(entity.DownloadFolder)
	controller.CreateFolder(entity.ArchiveFolder)
	a := api.Init()
	v1 := a.Group("/api/v1")
	api.RegisterUserHandlers(v1, *controller)
	a.Run(":8080")

}
