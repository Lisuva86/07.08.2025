package controller

import (
	"fmt"
	"net/http"
	"os"
	"zip_archive/entity"
)

// создание папки куда скачиваем
func (c *Controller) CreateFolder() error {
	folder := "test_downloads"
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

// проверка доступности
func (c *Controller) AvailabilityCheckURLs(urls []string) ([]entity.URLResult, error) {
	for _, item := range urls {
		resp, err := http.Head(item)

	}
}

//проверка разрешенных файлов на основании ентити ЮРЛрезалт

//скачивание файлов
//архивация файлов имя выдать по таске, по результату сохранить путь в таске
