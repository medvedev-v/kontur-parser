// file: excel.go
package xmlconverter

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	config "github.com/medvedev-v/kontur-parser/pkg/config"
	"github.com/xuri/excelize/v2"
)

type Excel interface {
	Header() []string
	Title() []string
	TitleRows() []interface{}
	HeaderRows() [][]interface{}
}

type Products struct {
	Products []struct {
		Name       string  `json:"name"`
		Price      float64 `json:"price"`
		Count      float64 `json:"count"`
		Code       string  `json:"code"`
		ShtrihCode string  `json:"shtrihcode"`
	}
}

func (td *Products) Title() []string {
	return []string{"Товары"}
}

func (td *Products) TitleRows() []interface{} {
	return []interface{}{""}
}

func (td *Products) Header() []string {
	return []string{"Название", "Цена", "Количество", "Код", "Штрихкод"}
}

func (td *Products) HeaderRows() [][]interface{} {
	var row [][]interface{}
	for _, value := range td.Products {
		r := []interface{}{
			value.Name,
			value.Price,
			value.Count,
			value.Code,
			value.ShtrihCode,
		}
		row = append(row, r)
	}
	return row
}

func SaveToExcel(products *XML) error {
	var excelData Products

	for _, product := range products.СведТов {
		productName := product.НаимТов
		productPrice := product.СтТовУчНал / product.КолТов
		productCount := product.КолТов
		productCode := ""
		productShtrihCode := ""

		if product.ИнфПолФХЖ2 != nil {
			for _, v := range product.ИнфПолФХЖ2 {
				switch v.Идентиф {
				case "КодТов":
					productCode = v.Значен
				case "штрихкод":
					productShtrihCode = v.Значен
				}
			}
		}

		productData := struct {
			Name       string  `json:"name"`
			Price      float64 `json:"price"`
			Count      float64 `json:"count"`
			Code       string  `json:"code"`
			ShtrihCode string  `json:"shtrihcode"`
		}{
			Name:       productName,
			Price:      productPrice,
			Count:      productCount,
			Code:       productCode,
			ShtrihCode: productShtrihCode,
		}
		excelData.Products = append(excelData.Products, productData)
	}

	file, err := ExcelGenerator(&excelData)
	if err != nil {
		return err
	}

	data, err := os.Create("output.xlsx")
	if err != nil {
		return err
	}
	defer data.Close()

	_, err = file.WriteTo(data)
	if err != nil {
		return err
	}

	return nil
}

func ExcelGenerator(data Excel) (*bytes.Buffer, error) {
	// Создаём новый файл
	file := excelize.NewFile()
	// Название листа
	sheetName := "Sheet1"

	// Получение заголовков
	header := data.Header()

	// Получение строк заголовков
	rows := data.HeaderRows()

	// Получение заголовков подзаголовков
	titles := data.Title()

	// Получение строк подзаголовков
	titleRows := data.TitleRows()

	// Создание заголовков
	for i, h := range titles {
		// Получение имени столбца по его номеру (A, B, C, ...)
		colName, _ := excelize.ColumnNumberToName(1)
		// Установка значения заголовка в ячейку
		file.SetCellValue(sheetName, colName+strconv.Itoa(i+1), h)
	}

	// Заполнение заголовков
	for i, h := range titleRows {
		// Получение имени столбца по его номеру (A, B, C, ...)
		colName, _ := excelize.ColumnNumberToName(2)
		// Установка значения заголовка в ячейку
		file.SetCellValue(sheetName, colName+strconv.Itoa(i+1), h)
	}

	// Создание заголовков колонок
	for i, h := range header {
		// Получение имени столбца по его номеру (A, B, C, ...)
		colName, _ := excelize.ColumnNumberToName(i + 1)
		// Установка значения заголовка колонки в ячейку
		file.SetCellValue(sheetName, colName+"7", h)
	}

	// Заполнение колонок
	for r, row := range rows {
		for c, val := range row {
			// Получение имени столбца по его номеру (A, B, C, ...)
			colName, _ := excelize.ColumnNumberToName(c + 1)
			// Установка значения в соответствующую ячейку
			file.SetCellValue(sheetName, colName+strconv.Itoa(r+8), val)
		}
	}

	// Запись файла в буфер
	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	// Запись файла на диск
	cfg := config.GetConfig()
	if err := file.SaveAs(cfg.DefaultOutputFilename); err != nil {
		fmt.Println(err)
	}

	// Возвращаем буфер и ошибку (если есть)
	return buf, nil
}
