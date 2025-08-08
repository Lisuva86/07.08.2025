package controller

import (
	"errors"
	"zip_archive/entity"
)

func (c *Controller) CreateTask() (int, *entity.Task, error) {
	taskBusyCount := 0
	for _, t := range c.storage {
		if t.Status == 1 {
			taskBusyCount++
		}
	}
	if taskBusyCount >= 3 {
		err := errors.New("server is busy")
		return -1, nil, err
	}
	var task entity.Task
	task.Status = 0
	task.ZipPath = ""
	task.URLSLice = make([]string, 0)
	c.storage = append(c.storage, task)
	return len(c.storage) - 1, &task, nil
}
func (c *Controller) GetTaskByID(id int) (*entity.Task, error){ 
	if id >= 0 && id < len(c.storage){
		return &c.storage[id], nil
	} else {
		err := errors.New("Task not found")
		return nil, err
	}
}
func (c *Controller) GetTaskStatusByID(id int) (int64,error){
	task, err := c.GetTaskByID(id)
	if err != nil{
		return -1,err
	}
	return task.Status, nil
}