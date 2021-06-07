package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	perfix := make([]int, len(nums)+1)
	perfix[0] = nums[0]
	for i := 1; i < len(nums)+1; i++ {
		perfix[i] = perfix[i-1] + nums[i-1]
	}
	hash := make(map[int]bool)
	for i := 2; i < len(perfix); i++ {
		hash[perfix[i-2]%k] = true
		_, found := hash[perfix[i]%k]
		if found {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23,2,6,4,7},6))
}
