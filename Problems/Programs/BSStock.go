package main

import "fmt"

func maxProfit(prices []int) (int, int, int) {
	if len(prices) == 0 {
		return 0, -1, -1
	}

	buy := prices[0]
	maxProfit := 0
	buyIndex := 0
	sellIndex := 0

	for i := 1; i < len(prices); i++ {
		// Checking for lower buy value
		if prices[i] < buy {
			buy = prices[i]
			buyIndex = i
		} else if prices[i]-buy > maxProfit { // Checking for higher profit
			maxProfit = prices[i] - buy
			sellIndex = i
		}
	}

	return maxProfit, buyIndex, sellIndex
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit, buyIndex, sellIndex := maxProfit(prices)
	fmt.Printf("Maximum profit: %d, Buy Index: %d, Sell Index: %d\n", maxProfit, buyIndex, sellIndex)
}
