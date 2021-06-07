package main
import "fmt"


func calc (prices []int) int{
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

func maxProfit_123(prices []int) int {
	max := 0
	for i:=0;i<len(prices);i++ {
		pointPorfit := calc(prices[:i]) + calc(prices[i:])
		if pointPorfit > max {
			max = pointPorfit
		}
	}
	return max
}




func main() {
	var arr = []int{1,2,4,2,5,7,2,4,9,0}
	fmt.Println(maxProfit_123(arr))
	arr = []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit_123(arr))
	arr = []int{1,2,3,4,5}
	fmt.Println(maxProfit_123(arr))
	arr = []int{7,6,4,3,1}
	fmt.Println(maxProfit_123(arr))
}
