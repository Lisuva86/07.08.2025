package api

import "github.com/gin-gonic/gin"

type API struct {
	*gin.Engine
}
type URIID struct {
	ID int `uri:"id" binding:"required"`
}