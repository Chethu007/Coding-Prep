package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"text/template"
)

func main() {
	// Define the XML template
	xmlTemplate := `
	<Request>
		<Amount>{{.Amount}}</Amount>
	   <Token>{{.Token}}</Token>
	   <SearchName>{{.SearchName}}</SearchName>
	</Request>
	`

	//xmlTemplate := `<Search><SearchItemLst><SearchItem><ItemName>{{.Token}}</ItemName><ItemVal>{{.Token}}</ItemVal></SearchItem><SearchItem><ItemName>CHECK NUMBER</ItemName><ItemVal>{{.SearchName}}</ItemVal></SearchItem><SearchItem><ItemName>AMOUNT</ItemName><ItemVal>{{.Token}}</ItemVal></SearchItem><SearchItem><ItemName>DIN</ItemName><ItemVal>{{.Token}}</ItemVal></SearchItem><SearchItem><ItemName>Process Date</ItemName><ItemVal>{{.Token}}</ItemVal></SearchItem></SearchItemLst></Search>`

	// Define the data structure
	amount := float64(75)
	data := struct {
		Token      string
		SearchName string
		Amount     string
	}{
		Token:      "your_token_here",
		SearchName: "your_search_name_here",
		Amount:     strconv.FormatFloat(amount, 'f', 4, 64), //fmt.Sprintf("%.4f", amount),
	}

	strconv.FormatFloat(amount, 'f', 4, 64)

	// Create a new template
	tmpl := template.New("request")

	// Parse the XML template
	tmpl, err := tmpl.Parse(xmlTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Buffer to store the XML output
	var xmlBuffer bytes.Buffer

	// Execute the template with the data and write the output to the buffer
	err = tmpl.Execute(&xmlBuffer, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Print the XML output
	fmt.Println("Generated XML:")
	fmt.Println(xmlBuffer.String())

	// Alternatively, if you want to encode the XML directly to a string
	xmlString, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	fmt.Println("Generated XML (using encoding/xml):")
	fmt.Println(string(xmlString))
}
