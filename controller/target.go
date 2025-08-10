package controller

import "zip_archive/entity"

func (c *Controller) AddTargetToTaskByTaskID(id int, URLs []string) (*entity.Task, error) {
	//проверить существует ли такая таска
	task, err := c.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	//проверить статус, ессли доне то выходим
	if task.Status == entity.TaskStatusDone {
		return task, nil
	}
	//проверить можем ли добавить задачу
	for _, item := range URLs {
		err = c.CheckTargetCountInTaskByID(id)
		if err != nil {
			//отправить на архивацию таску
			task, err = c.ArchiveTask(id)
			break
		}
		task.URLSLice = append(task.URLSLice, item)

	}

	return task, nil
}
