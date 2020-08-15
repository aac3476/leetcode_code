package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	for i:=1;i<len(prices);i++ {
		if prices[i] > prices[i-1] {
			res = res + prices[i] - prices[i-1]
		}
	}
	return res
}



func main() {
	var arr = []int{1,2,3,4,5}

	fmt.Println(maxProfit(arr))
}
