package test

import (
	"os"
	"testing"

	xmltoexcel "github.com/medvedev-v/kontur-parser/pkg/xmltoexcel"
)

func TestDefaultXMLParser_Parse(t *testing.T) {
	// Create test XML data
	testXML := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
	<xml>
		<СведТов НомСтр=\"1\" НаимТов=\"Товар 1\" ОКЕИ_Тов=\"796\" КолТов=\"2\" ЦенаТов=\"100\" СтТовБезНДС=\"200\" НалСт=\"20\" СтТовУчНал=\"240\" НаимЕдИзм=\"шт\">
			<ДопСведТов ПрТовРаб=\"Товар 1\" КодТов=\"00000001\" КрНаимСтрПр=\"Товар 1\" />
			<Акциз БезАкциз=\"Без акциза\" />
			<СумНал СумНал=\"40\" />
			<ИнфПолФХЖ2 Идентиф=\"КодТов\" Значен=\"00000001\" />
			<ИнфПолФХЖ2 Идентиф=\"штрихкод\" Значен=\"1234567890123\" />
		</СведТов>
	</xml>`

	// Create a new parser
	parser := &xmltoexcel.DefaultXMLParser{}

	// Parse the test XML
	result, err := parser.Parse([]byte(testXML))
	if err != nil {
		t.Fatalf("Failed to parse XML: %v", err)
	}

	// Check the results
	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	product := result.СведТов[0]
	if product.НаимТов != "Товар 1" {
		t.Errorf("Expected product name 'Товар 1', got '%s'", product.НаимТов)
	}

	if product.КолТов != 2 {
		t.Errorf("Expected quantity 2, got %f", product.КолТов)
	}

	if product.СтТовУчНал != 240 {
		t.Errorf("Expected total with tax 240, got %f", product.СтТовУчНал)
	}
}

func TestYashkinoXMLParser_Parse(t *testing.T) {
	// Create test XML data (same structure for now)
	testXML := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
	<xml>
		<СведТов НомСтр=\"1\" НаимТов=\"Товар 1\" ОКЕИ_Тов=\"796\" КолТов=\"2\" ЦенаТов=\"100\" СтТовБезНДС=\"200\" НалСт=\"20\" СтТовУчНал=\"240\" НаимЕдИзм=\"шт\">
			<ДопСведТов ПрТовРаб=\"Товар 1\" КодТов=\"00000001\" КрНаимСтрПр=\"Товар 1\" />
			<Акциз БезАкциз=\"Без акциза\" />
			<СумНал СумНал=\"40\" />
			<ИнфПолФХЖ2 Идентиф=\"КодТов\" Значен=\"00000001\" />
			<ИнфПолФХЖ2 Идентиф=\"штрихкод\" Значен=\"1234567890123\" />
		</СведТов>
	</xml>`

	// Create a new parser
	parser := &xmltoexcel.YashkinoXMLParser{}

	// Parse the test XML
	result, err := parser.Parse([]byte(testXML))
	if err != nil {
		t.Fatalf("Failed to parse XML: %v", err)
	}

	// Check the results
	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	product := result.СведТов[0]
	if product.НаимТов != "Товар 1" {
		t.Errorf("Expected product name 'Товар 1', got '%s'", product.НаимТов)
	}
}

func TestILSXMLParser_Parse(t *testing.T) {
	// Create test XML data
	testXML := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
	<xml>
		<СведТов НомСтр=\"1\" НаимТов=\"Товар 1\" ОКЕИ_Тов=\"796\" КолТов=\"2\" ЦенаТов=\"100\" СтТовБезНДС=\"200\" НалСт=\"20\" СтТовУчНал=\"240\" НаимЕдИзм=\"шт\">
			<ДопСведТов ПрТовРаб=\"Товар 1\" КодТов=\"00000001\" КолВедМарк=\"1\" />
			<Акциз БезАкциз=\"Без акциза\" />
			<СумНал СумНал=\"40\" />
		</СведТов>
	</xml>`

	// Create a new parser
	parser := &xmltoexcel.ILSXMLParser{}

	// Parse the test XML
	result, err := parser.Parse([]byte(testXML))
	if err != nil {
		t.Fatalf("Failed to parse XML: %v", err)
	}

	// Check the results
	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	product := result.СведТов[0]
	if product.НаимТов != "Товар 1" {
		t.Errorf("Expected product name 'Товар 1', got '%s'", product.НаимТов)
	}
}

func TestLoadFromFileWithSupplier(t *testing.T) {
	// Create a temporary test file
	testFile := "test_input.xml"
	testXML := `<?xml version=\"1.0\" encoding=\"UTF-8\"?>
	<xml>
		<СведТов НомСтр=\"1\" НаимТов=\"Товар 1\" ОКЕИ_Тов=\"796\" КолТов=\"2\" ЦенаТов=\"100\" СтТовБезНДС=\"200\" НалСт=\"20\" СтТовУчНал=\"240\" НаимЕдИзм=\"шт\">
			<ДопСведТов ПрТовРаб=\"Товар 1\" КодТов=\"00000001\" КрНаимСтрПр=\"Товар 1\" />
			<Акциз БезАкциз=\"Без акциза\" />
			<СумНал СумНал=\"40\" />
			<ИнфПолФХЖ2 Идентиф=\"КодТов\" Значен=\"00000001\" />
			<ИнфПолФХЖ2 Идентиф=\"штрихкод\" Значен=\"1234567890123\" />
		</СведТов>
	</xml>`

	// Write test file
	err := os.WriteFile(testFile, []byte(testXML), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile) // Clean up

	// Test with standard supplier
	result, err := xmltoexcel.LoadFromFileWithSupplier(testFile, "стандарт")
	if err != nil {
		t.Fatalf("Failed to load with standard supplier: %v", err)
	}

	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	// Test with Yashkino supplier
	result, err = xmltoexcel.LoadFromFileWithSupplier(testFile, "яшкино")
	if err != nil {
		t.Fatalf("Failed to load with Yashkino supplier: %v", err)
	}

	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	// Test with ILS supplier
	result, err = xmltoexcel.LoadFromFileWithSupplier(testFile, "ИЛС")
	if err != nil {
		t.Fatalf("Failed to load with ILS supplier: %v", err)
	}

	if len(result.СведТов) != 1 {
		t.Errorf("Expected 1 product, got %d", len(result.СведТов))
	}

	// Test with unsupported supplier
	_, err = xmltoexcel.LoadFromFileWithSupplier(testFile, "unknown")
	if err == nil {
		t.Error("Expected error for unsupported supplier, got nil")
	}
}

func TestLoadFromFileWithSupplier_FileNotFound(t *testing.T) {
	// Test with non-existent file
	_, err := xmltoexcel.LoadFromFileWithSupplier("nonexistent.xml", "стандарт")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}
