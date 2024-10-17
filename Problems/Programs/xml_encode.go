package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// InnerXML represents the inner XML structure
type InnerXML struct {
	XMLName xml.Name `xml:"Lst"`
	Items   []Item   `xml:"SearchItem"`
}

// Item represents an item in the nested XML structure
type Item struct {
	XMLName  xml.Name `xml:"SearchItem"`
	ItemName string   `xml:"ItemName"`
	ItemVal  string   `xml:"ItemVal"`
}

func main() {
	// Original XML string
	originalXML := `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="xxx" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<GetItem xmlns="http://www.com">
			<Token>1234567</Token>
			<SearchName>CS</SearchName>
			<SearchType>SEARCHTYPE</SearchType>
			<SearchXML>{.string}</SearchXML>
		</GetItem>
	</soap:Body>
</soap:Envelope>`

	// Replacement XML string
	replacementXML := `<Lst>
	<SearchItem>
		<ItemName>NUMBER</ItemName>
		<ItemVal>1234</ItemVal>
	</SearchItem>
	<SearchItem>
		<ItemName>1 NUMBER</ItemName>
		<ItemVal>1</ItemVal>
	</SearchItem>
	<SearchItem>
		<ItemName>AMOUNT</ItemName>
		<ItemVal>74</ItemVal>
	</SearchItem>
</Lst>`

	// Encode the replacement XML
	replacementXML = strings.ReplaceAll(replacementXML, "<", "&lt;")
	replacementXML = strings.ReplaceAll(replacementXML, ">", "&gt;")

	// Replace {.string} with the encoded replacement XML
	originalXML = strings.Replace(originalXML, "{.string}", replacementXML, 1)

	// Print the modified XML
	fmt.Println(originalXML)
}
