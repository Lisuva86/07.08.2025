package controller

import "zip_archive/entity"

func (c *Controller) AddTargetToTaskByTaskID(id int, URLs []string) (*entity.Task, error) {
	//проверить существует ли такая таска
	task, err := c.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	//проверить статус, ессли доне то выходим
	if task.Status == entity.TaskStatusEmpty {
		for _, item := range URLs {
			var url entity.URLResult
			url.Allowed = false
			url.Availability = false
			url.Error = ""
			url.URL = item
			url.FilePath = ""
			url.FileType = ""
			task.URLSLice = append(task.URLSLice, url)
			err = c.CheckTargetCountInTaskByID(id)
			if err != nil {
				//отправить на архивацию таску
				err = c.ArchiveTask(task)
				break
			}
		}
	}
	//проверить можем ли добавить задачу
	

	return task, nil
}
