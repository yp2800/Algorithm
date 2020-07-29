package leetcode

// 暴力
func maxProfit(prices []int) int {
	var maxProfit = 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0
	for i:=0; i< len(prices);i++{
		if prices[i] < minPrice{
			minPrice = prices[i]
		} else if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
