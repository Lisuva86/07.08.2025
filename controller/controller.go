package controller

import (
	"errors"
	"sync"
)

func (c *Controller) CheckTargetCount(id int) error{ //TASK
	if len(c.storage[id].URLSLice) >= 3{
		err := errors.New("max url in task")
		return err
	}
	return nil
}
func (c *Controller) ArchiveTask(id int) error{
	task, err := c.GetTaskByID(id)
	if err != nil{
		return err
	}
	task.Status = 1
	
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		//todo АРХИВ go
		task.Status = 2
		task.ZipPath = "Pipa"
		wg.Done()
	}()
	wg.Wait()
	return nil
}
func (c *Controller) GetArchivePath(id int) (string, error){
	task, err := c.GetTaskByID(id)
	if err != nil{
		return "", err
	}
	return task.ZipPath, nil

}