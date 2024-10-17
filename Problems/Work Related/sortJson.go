package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"sort"
//	"strings"
//)
//
//type Resp struct {
//	Period string `json:"Period,omitempty"`
//	Year   string `json:"Year"`
//	Id     string `json:"FileId,omitempty"`
//}
//
//type ByPeriodDescription []Resp
//
//func (a ByPeriodDescription) Len() int      { return len(a) }
//func (a ByPeriodDescription) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
//func (a ByPeriodDescription) Less(i, j int) bool {
//	return extractDescription(a[i].Period) < extractDescription(a[j].Period)
//}
//
//func extractDescription(period string) string {
//	parts := strings.Split(period, " - ")
//	return parts[1]
//}
//
//func main() {
//	// Sample data
//	data := []byte(`[
//		{"Period": "1 - January", "Year": "2022", "FileId": "1"},
//		{"Period": "3 - March", "Year": "2022", "FileId": "2"},
//		{"Period": "2 - February", "Year": "2022", "FileId": "3"}
//	]`)
//
//	// Unmarshal JSON into slice of Resp structs
//	var resps []Resp
//	err := json.Unmarshal(data, &resps)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	// Sort the slice based on the description of the Period field
//	sort.Sort(ByPeriodDescription(resps))
//
//	// Print sorted data
//	for _, resp := range resps {
//		fmt.Printf("Period: %s, Year: %s, FileId: %s\n", resp.Period, resp.Year, resp.Id)
//	}
//}
