package xmlconverter

import (
    "encoding/xml"
    "os"
)

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

// LoadFromFile - загружает данные из XML файла в структуру
func LoadFromFile(filename string) (*XML, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var products XML
    decoder := xml.NewDecoder(file)
    err = decoder.Decode(&products)
    if err != nil {
        return nil, err
    }

    return &products, nil
}
