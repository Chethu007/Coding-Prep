package main

import (
	"encoding/xml"
	"strings"
)

type Details struct {
	StartDate string `xml:"StartDate"`
	EndDate   string `xml:"EndDate"`
}

func main() {
	Xml := "PERldGFpbHM+PEFjY291bnRJZD4xMjUzOTwvQWNjb3VudElkPjxTdGFydERhdGU+MTk3MC0wMS0wMTwvU3RhcnREYXRlPjxFbmREYXRlPjIwMjUtMDUtMDI8L0VuZERhdGU+PElzUHJlbG9nb25TZXNzaW9uPkZhbHNlPC9Jc1ByZWxvZ29uU2Vzc2lvbj48U2Vzc2lvbklkPmZtY3dodXJrc3kybXVwdnVsdHlldXRxeDwvU2Vzc2lvbklkPjwvRGV0YWlscz4="
	var details *Details
	err := xml.Unmarshal([]byte(Xml), details)
	if err != nil {
		println("Error")
	}
	details.StartDate = strings.Replace(details.StartDate, "/", "-", -1)
	details.EndDate = strings.Replace(details.EndDate, "/", "-", -1)

	println(details.StartDate)
	println(details.EndDate)
}
