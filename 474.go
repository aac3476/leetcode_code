package main

import "fmt"

func max(a,b int) int {
	if a>b{
		return a
	} else {
		return b
	}
}

func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0{
		return 0
	}
	dp := make([][]int,(m+1))
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int,n+1)
	}

	for i := 0; i < len(strs); i++ {
		zeros,ones:=0,0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '0' {
				zeros ++
			} else {
				ones ++
			}
		}
		for p:=m;p>zeros-1;p--{
			for q:=n;q>ones-1;q--{
				dp[p][q] = max(dp[p][q], 1+dp[p- zeros][q-ones])
			}
		}
	}
	return dp[m][n]

}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
