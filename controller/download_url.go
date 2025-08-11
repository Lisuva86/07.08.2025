package controller

import (
	"errors"
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
func (c *Controller) CheckAvailability(urls []string) ([]entity.URLResult, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	var results []entity.URLResult

	for _, url := range urls {
		result := entity.URLResult{URL: url}

		resp, err := client.Head(url)
		if err != nil {
			// Если HEAD не прошёл — пробуем GET
			resp, err = client.Get(url)
			if err != nil {
				result.Error = err.Error()
				results = append(results, result)
				continue
			}
			// У GET нужно закрыть тело
			defer resp.Body.Close()
		} else {
			// HEAD — тело пустое, но всё равно нужно закрыть
			defer resp.Body.Close()
		}

		// Сохраняем Content-Type
		result.FileType = resp.Header.Get("Content-Type")

		// Проверяем статус
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			result.Availability = true
		} else {
			result.Error = fmt.Sprintf("HTTP %d", resp.StatusCode)
		}

		results = append(results, result)
	}
	//проверка на то, что все не доступно

	var errResult error
	errResult = nil
	if len(results) == 0 {
		errResult = errors.New("no links available")
	}

	return results, errResult
}

// проверка разрешенных файлов на основании ентити ЮРЛрезалт
func (c *Controller) CheckFileType(urls []entity.URLResult) ([]entity.URLResult, error) {
	var updatedResults []entity.URLResult

	for _, url := range urls {
		// Копируем результат, чтобы не изменять "на лету"
		r := url
		r.Allowed = false
		// Проверка по FileType
		contentType := r.FileType
		if contentType != "" {
			// Убираем параметры после ';'
			if idx := strings.Index(contentType, ";"); idx != -1 {
				contentType = contentType[:idx]
			}
			contentType = strings.TrimSpace(strings.ToLower(contentType))

			if entity.AllowedMIMETypes[contentType] {
				r.Allowed = true
				updatedResults = append(updatedResults, r)
				continue
			}
		}

		// Если FileType не помог — проверяем по расширению URL
		ext := strings.ToLower(filepath.Ext(r.URL))
		if entity.AllowedExtensions[ext] {
			r.Allowed = true
			updatedResults = append(updatedResults, r)
			continue
		}

		// Если ни тип, ни расширение не подошли
		r.Error = "file type not allowed"
		updatedResults = append(updatedResults, r)
	}

	return updatedResults, nil
}

// скачивание файлов
func (c *Controller) DownloadAllowedFiles(results []entity.URLResult) []entity.URLResult {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	var updatedResults []entity.URLResult

	for _, result := range results {
		// Копируем результат
		r := result

		// Пропускаем, если не доступен или не разрешён
		if !r.Availability || !r.Allowed {
			updatedResults = append(updatedResults, r)
			continue
		}

		// Генерируем безопасное имя файла
		filename := c.sanitizeFilename(filepath.Base(r.URL))
		if filename == "" {
			r.Error = "invalid filename in URL"
			updatedResults = append(updatedResults, r)
			continue
		}

		// Полный путь: test_download/filename.ext
		filePath := filepath.Join(entity.DownloadFolder, filename)

		// Скачиваем
		err := c.downloadFile(client, r.URL, filePath)
		if err != nil {
			r.Error = "download failed: " + err.Error()
		}
		r.FilePath = filePath
		updatedResults = append(updatedResults, r)
	}

	return updatedResults
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
