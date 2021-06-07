package main

import (
	"fmt"
)


func findMaxLength(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	max := 0
	perfix := make([]int, len(nums)+1)
	perfix[0] = 0
	hashmap := make(map[int]int)
	for i:= 1;i<len(nums) + 1;i++{
		if nums[i - 1] == 0{
			perfix[i] = perfix[i-1] - 1
		} else {
			perfix[i] = perfix[i-1] + 1
		}

	}
	for i:= 0;i<len(perfix);i++{
		_,found := hashmap[perfix[i]]
		if !found {
			hashmap[perfix[i]] = i
		} else {
			if i  - hashmap[perfix[i]] > max{
				max = i  - hashmap[perfix[i]]
			}
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxLength([]int{1,1,0,1,0,0}))
}
