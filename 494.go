package main

import "fmt"
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	neg := (sum - target) / 2
	if neg < 0 || (sum - target) % 2 != 0 || target > sum {
		return 0
	}
	dp := make([][]int, len(nums) + 1)
	dp[0] = make([]int, sum + 1)
	dp[0][0]=1
	for i := 1; i < len(nums) + 1; i++ {
		dp[i] = make([]int, neg + 1)
		for j:=0;j<=neg;j++{
			if j>=nums[i-1]{
				dp[i][j]=dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][neg]
}


func main() {
	fmt.Println(findTargetSumWays([]int{10},-10))
}
