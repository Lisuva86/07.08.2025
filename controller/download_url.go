package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	"zip_archive/entity"
)

// проверка доступности
func (c *Controller) CheckAvailability(task *entity.Task) error {
	client := &http.Client{Timeout: 10 * time.Second}
	
		
	for i := range task.URLSLice {
			url := &task.URLSLice[i]
		resp, err := client.Head(url.URL)
		if err != nil {
			// Если HEAD не прошёл — пробуем GET
			resp, err = client.Get(url.URL)
			if err != nil {
				url.Error = err.Error()
				continue
			}
			// У GET нужно закрыть тело
			defer resp.Body.Close()
		} else {
			// HEAD — тело пустое, но всё равно нужно закрыть
			defer resp.Body.Close()
		}

		// Сохраняем Content-Type
		url.FileType = resp.Header.Get("Content-Type")

		// Проверяем статус
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			url.Availability = true
		} else {
			url.Error = fmt.Sprintf("HTTP %d", resp.StatusCode)
		}
	}

	return  nil
}

// проверка разрешенных файлов на основании ентити ЮРЛрезалт
func (c *Controller) CheckFileType(task *entity.Task) error {
	for i := range task.URLSLice {
		url := &task.URLSLice[i]
		// Копируем результат, чтобы не изменять "на лету"
		
		url.Allowed = false
		// Проверка по FileType
		contentType := url.FileType
		if contentType != "" {
			// Убираем параметры после ';'
			if idx := strings.Index(contentType, ";"); idx != -1 {
				contentType = contentType[:idx]
			}
			contentType = strings.TrimSpace(strings.ToLower(contentType))

			if entity.AllowedMIMETypes[contentType] {
				url.Allowed = true
				
				continue
			}
		}

		// Если FileType не помог — проверяем по расширению URL
		ext := strings.ToLower(filepath.Ext(url.URL))
		if entity.AllowedExtensions[ext] {
			url.Allowed = true
			
			continue
		}
		// Если ни тип, ни расширение не подошли
		url.Error = "file type not allowed"
	}

	return nil
}

// скачивание файлов
func (c *Controller) DownloadAllowedFiles(task *entity.Task)  {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	for i := range task.URLSLice {
		url := &task.URLSLice[i]

		// Пропускаем, если не доступен или не разрешён
		if !url.Availability || !url.Allowed {
			continue
		}

		// Генерируем безопасное имя файла
		filename := c.sanitizeFilename(filepath.Base(url.URL))
		if filename == "" {
			url.Error = "invalid filename in URL"
			continue
		}

		// Полный путь: test_download/filename.ext
		filePath := filepath.Join(entity.DownloadFolder, filename)

		// Скачиваем
		err := c.downloadFile(client, url.URL, filePath)
		if err != nil {
			url.Error = "download failed: " + err.Error()
		}
		url.FilePath = filePath
	}
}
func (c *Controller) downloadFile(client *http.Client, url, filepath string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Создаём файл
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("cannot create file: %v", err)
	}
	defer out.Close()
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return fmt.Errorf("failed to copy data to file: %w", err)
    }

	return nil
}
func (c *Controller) sanitizeFilename(filename string) string {
	if filename == "" {
		return ""
	}

	// Декодируем URL-имя (например, %20 → пробел)
	decoded, err := url.PathUnescape(filename)
	if err != nil {
		decoded = filename // если ошибка — используем как есть
	}

	// Оставляем только имя файла (убираем путь)
	base := filepath.Base(decoded)

	// Разделяем имя и расширение
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// Очищаем имя от недопустимых символов
	cleanName := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z',
			r >= 'A' && r <= 'Z',
			r >= '0' && r <= '9',
			r == '-', r == '_', r == ' ':
			return r
		default:
			return '_'
		}
	}, name)

	// Ограничиваем длину имени (чтобы не было проблем с ФС)
	if len(cleanName) > 100 {
		cleanName = cleanName[:100]
	}

	// Если имя пустое — используем placeholder
	if cleanName == "" {
		cleanName = "file"
	}

	return cleanName + ext
}

//архивация файлов имя выдать по таске, по результату сохранить путь в таске
