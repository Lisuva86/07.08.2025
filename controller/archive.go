package controller

import (
	"sync"
	"zip_archive/entity"
)

func (c *Controller) ArchiveTask(id int) (*entity.Task, error) {
	task, err := c.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	task.Status = 1
	//создание папки если нет
	err = c.CreateFolder()
	if err != nil{
		return nil, err
	}
	//скачиваем файлы по юрл
	
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//todo АРХИВ go
		task.Status = 2
		task.ZipPath = "Pipa"
		wg.Done()
	}()
	wg.Wait()
	return task, nil
}
func (c *Controller) GetArchivePath(id int) (string, error){
	task, err := c.GetTaskByID(id)
	if err != nil{
		return "", err
	}
	return task.ZipPath, nil
}