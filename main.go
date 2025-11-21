// file: main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	config "github.com/medvedev-v/kontur-parser/pkg/config"
	converter "github.com/medvedev-v/kontur-parser/pkg/xmltoexcel"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %v\n", err)
		// Можно установить значения по умолчанию здесь
	} else {
		fmt.Println("Конфигурация загружена успешно")
	}

	fmt.Println("=== Конвертер XML в Excel ===")
	fmt.Println()
	
	// Список доступных поставщиков
	availableSuppliers := []string{"стандарт", "яшкино"}
	
	fmt.Println("Доступные поставщики:")
	for i, supplier := range availableSuppliers {
		fmt.Printf("  %d. %s\n", i+1, supplier)
	}
	fmt.Println()
	
	// Чтение названия поставщика
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите название поставщика: ")
	supplier, _ := reader.ReadString('\n')
	supplier = strings.TrimSpace(supplier)
	
	filename := cfg.DefaultInputFilename
	
	fmt.Printf("Обрабатываю файл %s для поставщика '%s'...\n", filename, supplier)
	
	// Загружаем данные
	products, err := converter.LoadFromFileWithSupplier(filename, supplier)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		fmt.Println("Нажмите Enter для закрытия...")
		reader.ReadString('\n')
		return
	}
	
	// Сохраняем в Excel
	err = converter.SaveToExcel(products)
	if err != nil {
		fmt.Printf("Ошибка сохранения в Excel: %v\n", err)
		fmt.Println("Нажмите Enter для закрытия...")
		reader.ReadString('\n')
		return
	}
	
	fmt.Println("✅ Парсинг успешно завершен!")
	fmt.Println("✅ Файл output.xlsx создан!")
	fmt.Println()
	fmt.Println("Нажмите Enter для закрытия...")
	reader.ReadString('\n')
}
