package gokatas

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		sv := prices[i] - minPrice
		if sv > maxProfit {
			maxProfit = sv
		}
	}

	return maxProfit
}
