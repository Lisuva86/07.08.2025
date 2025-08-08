package api

import (
	"zip_archive/entity"

	"github.com/gin-gonic/gin"
)
type createTaskResponce struct{
	Id int `json:"id"`
	Task entity.Task   `json:"task"`
}
func (h *handlers) postTaskHandler(c *gin.Context) {
	taskID,task, err := h.controller.CreateTask()
	if err != nil {
		c.JSON(429, gin.H{
			"error": err.Error(),
		})
		return
	}
	var responce createTaskResponce
	responce.Id = taskID
	responce.Task = *task
	c.JSON(201, &responce)
}
func (h *handlers) getTaskByIDHandler(c *gin.Context) {
	var uriID URIID
    err := c.ShouldBindUri(&uriID)
    if err != nil{
        c.JSON(429, gin.H{
			"error": err.Error(),
		})
    }
	task, err := h.controller.GetTaskByID(uriID.ID)
	if err != nil {
		c.JSON(429, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(201, &task)
}