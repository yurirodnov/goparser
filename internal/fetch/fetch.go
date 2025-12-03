package fetch

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// FetchURL получает HTML по заданному URL
func FetchURL(url string) ([]byte, error) {
	// Убираем лишние пробелы (у тебя в URL есть пробелы!)
	url = strings.TrimSpace(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("статус HTTP: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения тела: %w", err)
	}

	return body, nil
}