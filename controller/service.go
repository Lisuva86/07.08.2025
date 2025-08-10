package controller

import (
	"errors"
	"fmt"
	"os"
)

func (c *Controller) CheckTargetCountInTaskByID(id int) error { //TASK
	if len(c.storage[id].URLSLice) >= 3 {
		err := errors.New("max url in task")
		return err
	}
	return nil
}
func (c *Controller) CreateFolder(name string) error {
	folder := name
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Другая ошибка (например, доступ)
		return fmt.Errorf("ошибка при проверке папки %s: %v", folder, err)
	} else {
		fmt.Printf("Папка '%s' уже существует\n", folder)
	}

	return nil
}
