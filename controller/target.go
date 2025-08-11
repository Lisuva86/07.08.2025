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
			task.URLSLice = append(task.URLSLice, item)
			err = c.CheckTargetCountInTaskByID(id)
			if err != nil {
				//отправить на архивацию таску
				task, err = c.ArchiveTask(id)
				break
			}
		}
	}
	//проверить можем ли добавить задачу
	

	return task, nil
}
