package controller

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
	"zip_archive/entity"
)

func (c *Controller) ArchiveTask(id int) (*entity.Task, error) {
	task, err := c.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	task.Status = 1
	//чкачать в папку downloads
	// //создание папки если нет
	// err = c.CreateFolder()
	// if err != nil{
	// 	return nil, err
	// }
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
func (c *Controller) GetArchivePath(id int) (string, error) {
	task, err := c.GetTaskByID(id)
	if err != nil {
		return "", err
	}
	return task.ZipPath, nil
}
func (c *Controller) ArchiveFiles(files []entity.URLResult) (string, error) {
	// Генерируем имя архива с меткой времени
	timestamp := time.Now().Format("20060102_150405")
	archiveName := "archive_" + timestamp + ".zip"
	archivePath := filepath.Join(entity.ArchiveFolder, archiveName)

	// Создаём архив
	zipFile, err := os.Create(archivePath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Добавляем каждый файл
	for _, file := range files {
		filePath := file.FilePath
		// Проверяем, существует ли файл
		if filePath == "" {
			continue
		}
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			continue
		}

		// Открываем файл
		file, err := os.Open(filePath)
		if err != nil {
			continue // попробуем следующие
		}

		// Получаем имя файла (без пути)
		filename := filepath.Base(filePath)

		// Создаём запись в архиве
		zipEntry, err := zipWriter.Create(filename)
		if err != nil {
			file.Close()
			continue
		}

		// Копируем содержимое
		_, err = io.Copy(zipEntry, file)
		file.Close()

		if err != nil {
			// При ошибке копирования — прерываем, удаляем архив
			_ = zipWriter.Close()
			_ = zipFile.Close()
			_ = os.Remove(archivePath)
			return "", err
		}
	}

	return archivePath, nil // путь к созданному архиву
}
