package main

import (
	converter "github.com/medvedev-v/kontur-parser/pkg/xmltoexcel"
	"log"
)

func main() {
	products, err := converter.LoadFromFile("invoice.xml")
	if err != nil {
		log.Fatal(err)
	}

	converter.SaveToExcel(products)
}
