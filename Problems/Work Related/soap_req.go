package main

import (
	"encoding/xml"
	"fmt"
)

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XmlNS   string   `xml:"xmlns:soap,attr"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name      `xml:"soap:Header"`
	Headers []interface{} `xml:",any"` // Directly embed headers without extra <Headers> tag
}

type SOAPBody struct {
	XMLName xml.Name    `xml:"soap:Body"`
	Content interface{} `xml:",omitempty"`
}

// Predefined header types
type AuthToken struct {
	XMLName xml.Name `xml:"AuthToken"`
	Token   string   `xml:",chardata"`
}

type RequestID struct {
	XMLName xml.Name `xml:"RequestID"`
	ID      string   `xml:",chardata"`
}

// Custom header example
type CustomHeader struct {
	XMLName xml.Name `xml:"CustomHeader"`
	Data    string   `xml:",chardata"`
}

func main() {
	// Initialize SOAPEnvelope and SOAPHeader
	soapEnvelope := SOAPEnvelope{
		XmlNS: "http://schemas.xmlsoap.org/soap/envelope/",
	}

	// Initialize SOAPHeader and add predefined headers
	soapHeader := SOAPHeader{}
	soapHeader.Headers = append(soapHeader.Headers, AuthToken{Token: "12345"})
	soapHeader.Headers = append(soapHeader.Headers, RequestID{ID: "67890"})

	// Add a custom header
	customHeader := CustomHeader{
		Data: "Custom Data",
	}
	soapHeader.Headers = append(soapHeader.Headers, customHeader)

	// Assign the header to the SOAPEnvelope
	soapEnvelope.Header = &soapHeader

	// Print soapHeader.Headers after appending everything
	fmt.Println("soapHeader.Headers:")
	for i, header := range soapHeader.Headers {
		fmt.Printf("Header %d: %+v\n", i+1, header)
	}

	// Marshal SOAPEnvelope to XML and print the result
	output, err := xml.MarshalIndent(soapEnvelope, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling SOAPEnvelope:", err)
		return
	}

	fmt.Println("\nGenerated SOAP Request:")
	fmt.Println(string(output))
}
