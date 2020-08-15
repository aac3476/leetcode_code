package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min,max := prices[0],0
	for i:=1;i<len(prices);i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i] - min > max {
			max = prices[i] - min
		}
	}
	return max
}



func main() {
	var arr = []int{7,6,4,3,1}

	fmt.Println(maxProfit(arr))
}
