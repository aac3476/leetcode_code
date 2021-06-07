package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	maxwidth := len(triangle[height-1])
	if height == 1{
		return triangle[0][0]
	}
	for i:=1;i<height;i++ {
		for j:=0;j<len(triangle[i]);j++{
			if j==0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				continue
			}
			if j==len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][len(triangle[i-1])-1] + triangle[i][j]
				continue
			}
			triangle[i][j]=triangle[i][j] + min(triangle[i-1][j],triangle[i-1][j-1])
		}
	}
	result := triangle[height-1][maxwidth-1]
	for i:=0;i<maxwidth;i++{
		if triangle[height-1][i] < result {
			result = triangle[height-1][i]
		}
	}
	return result
}

func main(){
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
}