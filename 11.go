package main

import "fmt"

/*
盛水问题，双指针
 */

func maxArea(height []int) int {
	l,r,max := 0,len(height)-1,0
	for l<r {
		val := (r-l) * min(height[l],height[r])
		if val>max{
			max = val
		}
		if height[r] > height[l] {
			l++
		}else{
			r--
		}
	}
	return max
}




func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
