package main

import (
	"log"
	"os"

	"github.com/yurirodnov/goparser/internal/fetch"
)


func main(){

	// Можно взять URL из аргумента командной строки или захардкодить
	url := "https://ru.wikipedia.org/wiki/Белобрюхий_султаноносный_голубь"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	log.Printf("Скачивание: %s", url)

	data, err := fetch.FetchURL(url)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	log.Printf("Успешно! Получено %d байт", len(data))
	// Выводим начало (например, первые 300 символов)
	if len(data) > 300 {
		log.Println(string(data[:300]) + "...")
	} else {
		log.Println(string(data))
	}
}