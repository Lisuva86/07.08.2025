package controller

import (
	"errors"
)

func (c *Controller) CheckTargetCountInTaskByID(id int) error{ //TASK
	if len(c.storage[id].URLSLice) >= 3{
		err := errors.New("max url in task")
		return err
	}
	return nil
}

