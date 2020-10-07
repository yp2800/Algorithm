package leetcode

func maxProfit(prices []int) int {
	return calculate(prices, 0)
}

func calculate(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}
	max := 0
	for start := s; start < len(prices); start++ {
		maxProfit := 0
		for i := start + 1; i < len(prices); i++ {
			if prices[start] < prices[i] {
				profit := calculate(prices, i+1) + prices[i] - prices[start]
				if profit > maxProfit {
					maxProfit = profit
				}
			}
		}
		if maxProfit > max {
			max = maxProfit
		}
	}
	return max
}

func calculate2(prices []int) int {
	i := 0
	valley := prices[0]
	peak := prices[0]
	maxProfit := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxProfit += peak - valley
	}
	return maxProfit
}

func calculate3(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
