package controller

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
	"zip_archive/entity"
)

func (c *Controller) ArchiveTask(task *entity.Task) error {
	task.Status = 1
	//time.Sleep(10000 * time.Millisecond)
	err := c.CreateZip(task)
	if err != nil{
		return err
	}
	task.Status = 2
	return nil
}
func (c *Controller) GetArchivePath(id int) (string, error) {
	task, err := c.GetTaskByID(id)
	if err != nil {
		return "", err
	}
	return task.ZipPath, nil
}
func (c *Controller) ArchiveFiles(task *entity.Task) error {
	// Генерируем имя архива с меткой времени
	timestamp := time.Now().Format("20060102_150405")
	archiveName := task.TaskName + "_archive_" + timestamp + ".zip"
	archivePath := filepath.Join(entity.ArchiveFolder, archiveName)

	// Создаём архив
	zipFile, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Добавляем каждый файл

	for _, file := range task.URLSLice {
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
			return err
		}
	}
	task.ZipPath = archivePath

	return nil // путь к созданному архиву
}
func (c *Controller) CreateZip(task *entity.Task) error{
	err := c.CheckAvailability(task)
	if err != nil{
		return err
	}
	err = c.CheckFileType(task)
	if err != nil{
		return err
	}
	c.DownloadAllowedFiles(task)
	err = c.ArchiveFiles(task)
	if err != nil{
		return err
	}
	return nil

}
