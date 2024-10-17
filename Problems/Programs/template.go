package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

const startSavedSearchTemplate string = `                                                                                          {{- template "soap_header" -}}
<StartSavedSearch xmlns="http://www.profitstars.com/Synergy/SIMNET_Session/V1" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">{{- /**/ -}}
	<SessionToken>{{ .Token }}</SessionToken>                                                                                      {{- /**/ -}}
	<HitsPerPage>0</HitsPerPage>                                                                                                   {{- /**/ -}}
	<SearchName>{{ .SearchName }}</SearchName>                                                                                     {{- /**/ -}}
	<SavedSearchType>SIM_SEARCHTYPE_SYSTEM</SavedSearchType>                                                                       {{- /**/ -}}
	<SearchParamXML>{{ .CycleSearchXml }}</SearchParamXML>                                                                         {{- /**/ -}}
</StartSavedSearch>                                                                                                                {{- template "soap_footer" -}}`

const cycleSearchTemplate string = `           {{- /**/ -}}
<Search>                                       {{- /**/ -}}
	<SearchItemLst>                            {{- /**/ -}}
		<SearchItem>                           {{- /**/ -}}
			<ItemName>ACCOUNT NUMBER</ItemName>{{- /**/ -}}
			<Compare>=</Compare>               {{- /**/ -}}
			<ItemVal>{{ .AcctNum }}</ItemVal>  {{- /**/ -}}
			<Connector>&amp;</Connector>       {{- /**/ -}}
		</SearchItem>                          {{- /**/ -}}
		<SearchItem>                           {{- /**/ -}}
			<ItemName>DOC DATE</ItemName>      {{- /**/ -}}
			<Compare>&gt;=</Compare>           {{- /**/ -}}
			<ItemVal>{{ .Period }}</ItemVal>   {{- /**/ -}}
		</SearchItem>                          {{- /**/ -}}
	</SearchItemLst>                           {{- /**/ -}}
</Search>                                      {{- /**/ -}}`

//new1 := `<Search>
//	<SearchItemLst>
//		<SearchItem>
//			<ItemName>{{.ItemName}}</ItemName>
//			<ItemVal>{{.AccountNumber}}</ItemVal>
//		</SearchItem>
//		<SearchItem>
//			<ItemName>CHECK NUMBER</ItemName>
//			<ItemVal>{{.CheckNumber}}</ItemVal>
//		</SearchItem>
//		<SearchItem>
//			<ItemName>AMOUNT</ItemName>
//			<ItemVal>{{.Amount}}</ItemVal>
//		</SearchItem>
//		<SearchItem>
//			<ItemName>DIN</ItemName>
//			<ItemVal>{{.TraceNumber}}</ItemVal>
//		</SearchItem>
//		<SearchItem>
//			<ItemName>Process Date</ItemName>
//			<ItemVal>{{.ProcessDate}}</ItemVal>
//		</SearchItem>
//	</SearchItemLst>
//</Search>`

func main() {
	// Define the XML template
	//	xmlTemplate := `
	//<Request>
	//    <Token>{{.Token}}</Token>
	//    <SearchName>{{.SearchName}}</SearchName>
	//</Request>
	//`

	//	xmlTemplate1 := `<?xml version="1.0" encoding="utf-8"?>
	//<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//    <soap:Body>
	//        <GetItemBySearch xmlns="http://www.profitstars.com/Synergy/SIMNET_Session/V1">
	//            <SessionToken>{{.Token}}</SessionToken>
	//            <SearchName>{{.SearchName}}</SearchName>
	//            <SavedSearchType>{{.SearchType}}</SavedSearchType>
	//            <SearchParamXML>{{.SearchXml}}</SearchParamXML>
	//            <SearchTimeout>{{.Timeout}}</SearchTimeout>
	//            <SynItemRequestObj>
	//                <Options>
	//                    <SynOptionEnums>{{.OptionEnum}}</SynOptionEnums>
	//                </Options>
	//                <StartPage>{{.Startpage}}</StartPage>
	//                <NumberOfPages>{{.NumOfPages}}</NumberOfPages>
	//                <Extension>{{.Extension}}</Extension>
	//            </SynItemRequestObj>
	//        </GetItemBySearch>
	//    </soap:Body>
	//</soap:Envelope>`

	xmlTemplate := template.New("Image_request")
	template.Must(xmlTemplate.Parse(SoapHeaderTemplate))
	template.Must(xmlTemplate.Parse(startSavedSearchTemplate))
	template.Must(xmlTemplate.Parse(SoapFooterTemplate))
	// Define the data structure
	data := struct {
		Token          string
		SearchName     string
		SearchType     string
		SearchXml      string
		Timeout        int64
		OptionEnum     string
		Startpage      int64
		NumOfPages     int64
		Extension      string
		CycleSearchXml string
	}{
		Token:          "token",
		SearchName:     "searchname",
		SearchType:     "searchtype",
		SearchXml:      "cycleSearchTemplate",
		Timeout:        10,
		OptionEnum:     "",
		Startpage:      1,
		NumOfPages:     1,
		Extension:      "extension",
		CycleSearchXml: "cycleSearchTemplate",
	}

	// Create a new template
	//tmpl := template.New("Image_request")

	// Parse the XML template
	var tpl bytes.Buffer
	err := xmlTemplate.Execute(&tpl, data)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	xmlBytes, err := xml.MarshalIndent(tpl.String(), "", "    ")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return
	}
	fmt.Println(tpl.String())
	println(string(xmlBytes))

	// Buffer to store the XML output
	//var xmlBuffer bytes.Buffer
	//
	//// Execute the template with the data and write the output to the buffer
	//err = tmpl.Execute(&xmlBuffer, data)
	//if err != nil {
	//	fmt.Println("Error executing template:", err)
	//	return
	//}
	//
	//// Print the XML output
	//fmt.Println("Generated XML:")
	//fmt.Println(xmlBuffer.String())
	//
	//// Alternatively, if you want to encode the XML directly to a string
	//xmlString, err := xml.MarshalIndent(data, "", "    ")
	//if err != nil {
	//	fmt.Println("Error marshalling XML:", err)
	//	return
	//}
	//
	//fmt.Println("Generated XML (using encoding/xml):")
	//fmt.Println(string(xmlString))
}

func MakeRequest[T any](data T, templateData, templateName string) (string, error) {
	generic := template.New(templateName)

	template.Must(generic.Parse(SoapHeaderTemplate))
	template.Must(generic.Parse(SoapFooterTemplate))
	template.Must(generic.Parse(templateData))

	out := bytes.NewBuffer(make([]byte, 0, len(templateData)))
	if err := generic.Execute(out, data); err != nil {
		return "", err
	}

	return out.String(), nil
}

const SoapHeaderTemplate string = `{{- define "soap_header" -}}
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">{{- /**/ -}}
	<s:Body>{{- end -}}`

const SoapFooterTemplate string = `{{- define "soap_footer" -}}
	</s:Body>{{- /**/ -}}
</s:Envelope>{{- end -}}`
