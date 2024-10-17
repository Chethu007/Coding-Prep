package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Resp struct {
	Period string `json:"Period,omitempty"`
	Year   string `json:"Year"`
	Id     string `json:"FileId,omitempty"`
}

// Custom sorting function for sorting by description, year, and month of Period
func sortByDescriptionAndDate(respSlice []*Resp) {
	// Define a custom sorting function
	sort.Slice(respSlice, func(i, j int) bool {
		// Extract description, year, and month from Period
		descI, yearI, monthI := extractDescYearMonth(respSlice[i].Period)
		descJ, yearJ, monthJ := extractDescYearMonth(respSlice[j].Period)

		// Compare descriptions
		if descI != descJ {
			return descI < descJ
		}
		// If descriptions are equal, compare year
		if yearI != yearJ {
			return yearI > yearJ // Sorting in descending order of year
		}
		// If years are equal, compare month
		return monthI > monthJ // Sorting in descending order of month
	})
}

// Helper function to extract description, year, and month from Period
func extractDescYearMonth(period string) (string, int, int) {
	parts := strings.Split(period, " - ")
	desc := parts[2]

	yearMonth := strings.Split(parts[0], " ")
	year, _ := strconv.Atoi(yearMonth[0])
	month, _ := strconv.Atoi(yearMonth[1])
	return desc, year, month
}

func main() {
	// Sample data
	resps := []*Resp{
		{"2022 1 January", "2022", "1"},
		{"2022 3 March", "2022", "2"},
		{"2022 2 February", "2022", "3"},
		{"2023 1 January", "2023", "4"},
	}

	// Sort the slice based on the description, year, and month of the Period field
	sortByDescriptionAndDate(resps)

	// Print sorted data
	for _, resp := range resps {
		fmt.Printf("Period: %s, Year: %s, FileId: %s\n", resp.Period, resp.Year, resp.Id)
	}
}
