package api

import (
	"net/http"
	"zip_archive/entity"

	"github.com/gin-gonic/gin"
)

type getStatusResponce struct{
	Status entity.TaskStatus
	Task entity.Task
}
func (h *handlers) getTaskStatusByIDHandler(c *gin.Context) {
	var uriID URIID
	err := c.ShouldBindUri(&uriID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	//получили таску
	task, err := h.controller.GetTaskByID(uriID.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//получили статус
	status, err := h.controller.GetTaskStatusByID(uriID.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var responce getStatusResponce
	responce.Task = *task
	responce.Status = status
	c.JSON(http.StatusOK, &responce)
}