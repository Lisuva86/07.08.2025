package api

import (
	"net/http"
	"zip_archive/entity"

	"github.com/gin-gonic/gin"
)

type createTargetResponce struct {
	Task entity.Task `json:"task"`
}
func (h *handlers) postTargetToTaskHandler(c *gin.Context) {
    var uriID URIID
    err := c.ShouldBindUri(&uriID)
    if err != nil{
        c.JSON(429, gin.H{
			"error": err.Error(),
		})
    }
    var body entity.Target
    err = c.ShouldBindJSON(&body)
	if err != nil {
        c.JSON(429, gin.H{
			"error": err.Error(),
		})
		return
	}
    task, err := h.controller.AddTargetToTaskByTaskID(uriID.ID, body.URL)
    if err != nil{
        c.JSON(429, gin.H{
			"error": err.Error(),
		})
        return
    }
    var responce createTargetResponce
    responce.Task = *task

   
    c.JSON(http.StatusOK, responce)
}