package api

import (
	"zip_archive/controller"

	"github.com/gin-gonic/gin"
)

func Init() *API {
	router := gin.Default()

	return &API{router}
}

type handlers struct {
	controller controller.Controller
}

func New(controller controller.Controller) *handlers {
	return &handlers{
		controller: controller,
	}
}

func RegisterUserHandlers(routerGroup *gin.RouterGroup, controller controller.Controller) {
	h := New(controller)
	//--------------------------------------------------------------------------------task
	{
		tasks := routerGroup.Group("/tasks")
		tasks.POST("", h.postTaskHandler)
	}
	{
		task := routerGroup.Group("/task/:id")
		task.GET("", h.getTaskByIDHandler)
	}
	//--------------------------------------------------------------------------------target_url
	{
		target_url := routerGroup.Group("/target-to-task/:id")
		target_url.POST("", h.postTargetToTaskHandler)
		target_url.GET("", h.deleteAFTERTESTHANDLER)
	}
	//--------------------------------------------------------------------------------status
	{
		status := routerGroup.Group("/task-status/:id")
		status.GET("", h.getTaskStatusByIDHandler)
	}
}
