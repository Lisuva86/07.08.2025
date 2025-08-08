package controller

import "zip_archive/entity"

func (c *Controller) AddTargetToTaskByTaskID(id int, URLs []string) (*entity.Task, error) {
	//проверить существует ли такая таска
	task, err := c.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	//проверить можем ли добавить задачу
	for _, item := range URLs {
		err = c.CheckTargetCount(id)
		if err != nil {
			break
		}
		task.URLSLice = append(task.URLSLice, item)
	}

	return task, nil
}