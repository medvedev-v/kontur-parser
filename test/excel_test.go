package test

import (
	"os"
	"testing"

	xmltoexcel "github.com/medvedev-v/kontur-parser/pkg/xmltoexcel"
)

func TestProducts_Title(t *testing.T) {
	products := &xmltoexcel.Products{}
	result := products.Title()

	if len(result) != 1 {
		t.Errorf("Expected 1 title, got %d", len(result))
	}

	if result[0] != "Товары" {
		t.Errorf("Expected title 'Товары', got '%s'", result[0])
	}
}

func TestProducts_TitleRows(t *testing.T) {
	products := &xmltoexcel.Products{}
	result := products.TitleRows()

	if len(result) != 1 {
		t.Errorf("Expected 1 title row, got %d", len(result))
	}

	if result[0] != "" {
		t.Errorf("Expected empty string, got '%s'", result[0])
	}
}

func TestProducts_Header(t *testing.T) {
	products := &xmltoexcel.Products{}
	result := products.Header()

	expectedHeaders := []string{"Название", "Цена", "Количество", "Код", "Штрихкод"}

	if len(result) != len(expectedHeaders) {
		t.Errorf("Expected %d headers, got %d", len(expectedHeaders), len(result))
	}

	for i, header := range expectedHeaders {
		if result[i] != header {
			t.Errorf("Expected header '%s' at position %d, got '%s'", header, i, result[i])
		}
	}
}

func TestProducts_HeaderRows(t *testing.T) {
	// Create test data
	products := &xmltoexcel.Products{
		Products: []struct {
			Name       string  `json:"name"`
			Price      float64 `json:"price"`
			Count      float64 `json:"count"`
			Code       string  `json:"code"`
			ShtrihCode string  `json:"shtrihcode"`
		}{
			{
				Name:       "Товар 1",
				Price:      100.50,
				Count:      2.0,
				Code:       "00000001",
				ShtrihCode: "1234567890123",
			},
			{
				Name:       "Товар 2",
				Price:      50.25,
				Count:      5.0,
				Code:       "00000002",
				ShtrihCode: "9876543210987",
			},
		},
	}

	result := products.HeaderRows()

	if len(result) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(result))
	}

	// Check first row
	if len(result[0]) != 5 {
		t.Errorf("Expected 5 columns in first row, got %d", len(result[0]))
	}

	if result[0][0] != "Товар 1" {
		t.Errorf("Expected 'Товар 1' in first column, got '%v'", result[0][0])
	}

	if result[0][1] != 100.50 {
		t.Errorf("Expected 100.50 in second column, got '%v'", result[0][1])
	}

	if result[0][2] != 2.0 {
		t.Errorf("Expected 2.0 in third column, got '%v'", result[0][2])
	}

	if result[0][3] != "00000001" {
		t.Errorf("Expected '00000001' in fourth column, got '%v'", result[0][3])
	}

	if result[0][4] != "1234567890123" {
		t.Errorf("Expected '1234567890123' in fifth column, got '%v'", result[0][4])
	}

	// Check second row
	if len(result[1]) != 5 {
		t.Errorf("Expected 5 columns in second row, got %d", len(result[1]))
	}

	if result[1][0] != "Товар 2" {
		t.Errorf("Expected 'Товар 2' in first column, got '%v'", result[1][0])
	}

	if result[1][1] != 50.25 {
		t.Errorf("Expected 50.25 in second column, got '%v'", result[1][1])
	}

	if result[1][2] != 5.0 {
		t.Errorf("Expected 5.0 in third column, got '%v'", result[1][2])
	}

	if result[1][3] != "00000002" {
		t.Errorf("Expected '00000002' in fourth column, got '%v'", result[1][3])
	}

	if result[1][4] != "9876543210987" {
		t.Errorf("Expected '9876543210987' in fifth column, got '%v'", result[1][4])
	}
}

func TestSaveToExcel(t *testing.T) {
	// Create test XML data
	testXML := &xmltoexcel.XML{
		СведТов: []struct {
			Text        string  `xml:",chardata"`
			НомСтр      string  `xml:"НомСтр,attr"`
			НаимТов     string  `xml:"НаимТов,attr"`
			ОКЕИТов     string  `xml:"ОКЕИ_Тов,attr"`
			КолТов      float64 `xml:"КолТов,attr"`
			ЦенаТов     string  `xml:"ЦенаТов,attr"`
			СтТовБезНДС string  `xml:"СтТовБезНДС,attr"`
			НалСт       string  `xml:"НалСт,attr"`
			СтТовУчНал  float64 `xml:"СтТовУчНал,attr"`
			НаимЕдИзм   string  `xml:"НаимЕдИзм,attr"`
			ДопСведТов  struct {
				Text        string `xml:",chardata"`
				ПрТовРаб    string `xml:"ПрТовРаб,attr"`
				КодТов      string `xml:"КодТов,attr"`
				КрНаимСтрПр string `xml:"КрНаимСтрПр"`
			} `xml:"ДопСведТов"`
			Акциз struct {
				Text     string `xml:",chardata"`
				БезАкциз string `xml:"БезАкциз"`
			} `xml:"Акциз"`
			СумНал struct {
				Text   string `xml:",chardata"`
				СумНал string `xml:"СумНал"`
			} `xml:"СумНал"`
			ИнфПолФХЖ2 []struct {
				Text    string `xml:",chardata"`
				Идентиф string `xml:"Идентиф,attr"`
				Значен  string `xml:"Значен,attr"`
			} `xml:"ИнфПолФХЖ2"`
		}{
			{
				НаимТов:    "Товар 1",
				КолТов:     2.0,
				СтТовУчНал: 240.0,
				ДопСведТов: struct {
					Text        string `xml:",chardata"`
					ПрТовРаб    string `xml:"ПрТовРаб,attr"`
					КодТов      string `xml:"КодТов,attr"`
					КрНаимСтрПр string `xml:"КрНаимСтрПр"`
				}{
					КодТов: "00000001",
				},
				ИнфПолФХЖ2: []struct {
					Text    string `xml:",chardata"`
					Идентиф string `xml:"Идентиф,attr"`
					Значен  string `xml:"Значен,attr"`
				}{
					{
						Идентиф: "КодТов",
						Значен:  "00000001",
					},
					{
						Идентиф: "штрихкод",
						Значен:  "1234567890123",
					},
				},
			},
			{
				НаимТов:    "Товар 2",
				КолТов:     5.0,
				СтТовУчНал: 250.0,
				ДопСведТов: struct {
					Text        string `xml:",chardata"`
					ПрТовРаб    string `xml:"ПрТовРаб,attr"`
					КодТов      string `xml:"КодТов,attr"`
					КрНаимСтрПр string `xml:"КрНаимСтрПр"`
				}{
					КодТов: "00000002",
				},
				ИнфПолФХЖ2: []struct {
					Text    string `xml:",chardata"`
					Идентиф string `xml:"Идентиф,attr"`
					Значен  string `xml:"Значен,attr"`
				}{
					{
						Идентиф: "КодТов",
						Значен:  "00000002",
					},
					{
						Идентиф: "штрихкод",
						Значен:  "9876543210987",
					},
				},
			},
		},
	}

	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_output_*.xlsx")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Save the temporary file path
	originalFilename := "output.xlsx"

	// Override the original filename with our temp file
	err = os.Rename(tempFile.Name(), originalFilename)
	if err != nil {
		t.Fatalf("Failed to rename temp file: %v", err)
	}
	defer os.Rename(originalFilename, tempFile.Name()) // Restore original

	// Test SaveToExcel
	if err := xmltoexcel.SaveToExcel(testXML); err != nil {
		t.Fatalf("SaveToExcel failed: %v", err)
	}

	// Check if file was created
	if _, err := os.Stat(originalFilename); os.IsNotExist(err) {
		t.Error("Expected output.xlsx to be created, but it doesn't exist")
	}

	// Clean up
	os.Remove(originalFilename)
}

func TestExcelGenerator(t *testing.T) {
	// Test with empty data
	var emptyData xmltoexcel.Products
	buf, err := xmltoexcel.ExcelGenerator(&emptyData)
	if err != nil {
		t.Fatalf("Failed to generate Excel with empty data: %v", err)
	}

	if buf == nil {
		t.Fatal("Expected non-nil buffer with empty data, got nil")
	}

	// Test with nil data
	var nilData xmltoexcel.Products
	buf, err = xmltoexcel.ExcelGenerator(&nilData)
	if err != nil {
		t.Fatalf("Failed to generate Excel with nil data: %v", err)
	}

	if buf == nil {
		t.Fatal("Expected non-nil buffer with nil data, got nil")
	}

	// Test with actual data
	testData := &xmltoexcel.Products{
		Products: []struct {
			Name       string  `json:"name"`
			Price      float64 `json:"price"`
			Count      float64 `json:"count"`
			Code       string  `json:"code"`
			ShtrihCode string  `json:"shtrihcode"`
		}{
			{
				Name:       "Товар 1",
				Price:      100.50,
				Count:      2.0,
				Code:       "00000001",
				ShtrihCode: "1234567890123",
			},
		},
	}

	buf, err = xmltoexcel.ExcelGenerator(testData)
	if err != nil {
		t.Fatalf("Failed to generate Excel with test data: %v", err)
	}

	if buf == nil {
		t.Fatal("Expected non-nil buffer with test data, got nil")
	}

	if buf.Len() == 0 {
		t.Error("Expected non-empty buffer with test data, got empty buffer")
	}

	// Test that the buffer contains Excel data
	// We can't parse the Excel file here, but we can check that it's not obviously wrong
	// A valid XLSX file should have a certain structure, but we'll just check that it's not empty
	// and contains some expected strings in a real test environment
	// For now, we just verify that the function completes without error
}
