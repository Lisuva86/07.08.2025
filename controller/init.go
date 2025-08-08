package controller

import "zip_archive/entity"

type Controller struct {
	storage []entity.Task
}
func New() *Controller{
	return &Controller{storage: make([]entity.Task, 0)}
}