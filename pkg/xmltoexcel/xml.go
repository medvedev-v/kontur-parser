// file: xml.go
package xmlconverter

import (
    "encoding/xml"
    "os"
    "fmt"
    "strings"
)

// Основная структура XML (должна быть объявлена)
type XML struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	СведТов []struct {
		Text        string `xml:",chardata"`
		НомСтр      string `xml:"НомСтр,attr"`
		НаимТов     string `xml:"НаимТов,attr"`
		ОКЕИТов     string `xml:"ОКЕИ_Тов,attr"`
		КолТов      float64 `xml:"КолТов,attr"`
		ЦенаТов     string `xml:"ЦенаТов,attr"`
		СтТовБезНДС string `xml:"СтТовБезНДС,attr"`
		НалСт       string `xml:"НалСт,attr"`
		СтТовУчНал  float64 `xml:"СтТовУчНал,attr"`
		НаимЕдИзм   string `xml:"НаимЕдИзм,attr"`
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
	} `xml:"СведТов"`
}

// XMLParser интерфейс для различных парсеров XML
type XMLParser interface {
    Parse(xmlData []byte) (*XML, error)
}

// Структура для стандартного XML
type DefaultXMLParser struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	СведТов []struct {
		Text        string `xml:",chardata"`
		НомСтр      string `xml:"НомСтр,attr"`
		НаимТов     string `xml:"НаимТов,attr"`
		ОКЕИТов     string `xml:"ОКЕИ_Тов,attr"`
		КолТов      float64 `xml:"КолТов,attr"`
		ЦенаТов     string `xml:"ЦенаТов,attr"`
		СтТовБезНДС string `xml:"СтТовБезНДС,attr"`
		НалСт       string `xml:"НалСт,attr"`
		СтТовУчНал  float64 `xml:"СтТовУчНал,attr"`
		НаимЕдИзм   string `xml:"НаимЕдИзм,attr"`
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
	} `xml:"СведТов"`
}

// Parse реализация метода Parse для стандартного XML
func (p *DefaultXMLParser) Parse(xmlData []byte) (*XML, error) {
    var products XML
    err := xml.Unmarshal(xmlData, &products)
    if err != nil {
        return nil, err
    }
    return &products, nil
}

// Структура для Яшкино XML
type YashkinoXMLParser struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	СведТов []struct {
		Text        string `xml:",chardata"`
		НомСтр      string `xml:"НомСтр,attr"`
		НаимТов     string `xml:"НаимТов,attr"`
		ОКЕИТов     string `xml:"ОКЕИ_Тов,attr"`
		КолТов      float64 `xml:"КолТов,attr"`
		ЦенаТов     string `xml:"ЦенаТов,attr"`
		СтТовБезНДС string `xml:"СтТовБезНДС,attr"`
		НалСт       string `xml:"НалСт,attr"`
		СтТовУчНал  float64 `xml:"СтТовУчНал,attr"`
		НаимЕдИзм   string `xml:"НаимЕдИзм,attr"`
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
	} `xml:"СведТов"`
}

// Parse реализация метода Parse для Яшкино XML
func (p *YashkinoXMLParser) Parse(xmlData []byte) (*XML, error) {
    // Пока используем стандартный парсер как заглушку
    // В будущем здесь будет специальная логика для XML от Яшкино
    var products XML
    err := xml.Unmarshal(xmlData, &products)
    if err != nil {
        return nil, err
    }
    return &products, nil
}

// LoadFromFileWithSupplier загружает XML с указанием поставщика
func LoadFromFileWithSupplier(filename string, supplier string) (*XML, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("не удалось прочитать файл %s: %v", filename, err)
    }
    
    // Нормализуем название поставщика
    supplier = strings.ToLower(strings.TrimSpace(supplier))
    
    // Выбираем парсер в зависимости от поставщика
    var parser XMLParser
    
    switch supplier {
    case "стандарт", "default":
        parser = &DefaultXMLParser{}
    case "яшкино", "yashkino":
        parser = &YashkinoXMLParser{}
	case "ИЛС", "ILS", "илс":
        parser = &ILSXMLParser{}	
    default:
        return nil, fmt.Errorf("поставщик '%s' не поддерживается", supplier)
    }
    
    return parser.Parse(data)
}

// Структура для ИЛС XML
type ILSXMLParser struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	СведТов []struct {
		Text        string `xml:",chardata"`
		НомСтр      string `xml:"НомСтр,attr"`
		НаимТов     string `xml:"НаимТов,attr"`
		ОКЕИТов     string `xml:"ОКЕИ_Тов,attr"`
		КолТов      float64 `xml:"КолТов,attr"`
		ЦенаТов     string `xml:"ЦенаТов,attr"`
		СтТовБезНДС string `xml:"СтТовБезНДС,attr"`
		НалСт       string `xml:"НалСт,attr"`
		СтТовУчНал  float64 `xml:"СтТовУчНал,attr"`
		НаимЕдИзм   string `xml:"НаимЕдИзм,attr"`
		ДопСведТов  struct {
			Text        string `xml:",chardata"`
			ПрТовРаб    string `xml:"ПрТовРаб,attr"`
			КодТов      string `xml:"КодТов,attr"`
			КолВедМарк string `xml:"КолВедМарк,attr"`
		} `xml:"ДопСведТов"`
		Акциз struct {
			Text     string `xml:",chardata"`
			БезАкциз string `xml:"БезАкциз"`
		} `xml:"Акциз"`
		СумНал struct {
			Text   string `xml:",chardata"`
			СумНал string `xml:"СумНал"`
		} `xml:"СумНал"`
	} `xml:"СведТов"`
}

// Parse реализация метода Parse для ILS XML
func (p *ILSXMLParser) Parse(xmlData []byte) (*XML, error) {
    var products XML
    err := xml.Unmarshal(xmlData, &products)
    if err != nil {
        return nil, err
    }
    return &products, nil
}
