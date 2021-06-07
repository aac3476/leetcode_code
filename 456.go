package main

import "fmt"

func find132pattern(nums []int) bool {
	leftMinArray := make([]int,len(nums))
	if len(nums) < 3 {
		return false
	}
	leftMinArray[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i - 1] < leftMinArray[i - 1] {
			leftMinArray[i] = nums[i - 1]
		} else {
			leftMinArray[i] = leftMinArray[i - 1]
		}
	}
	for i := 1; i < len(nums)-1; i++ {
		for j:=i+1;j<len(nums);j++{
			if nums[j] > leftMinArray[i] && nums[j] < nums[i]{
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{3, 5, 0, 3, 4}
	fmt.Println(find132pattern(arr))
	arr = []int{1,2,3,4}
	fmt.Println(find132pattern(arr))
	arr = []int{3,1,4,2}
	fmt.Println(find132pattern(arr))
	arr = []int{-1,3,2,0}
	fmt.Println(find132pattern(arr))
}
