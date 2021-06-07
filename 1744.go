package main

import "fmt"

func canEat(candiesCount []int, queries [][]int) []bool {
	if len(candiesCount) < 1{
		return []bool{false}
	}
	perfix := make([]int,len(candiesCount) + 1)
	perfix[1] = candiesCount[0]
	for i:=2;i<len(candiesCount) + 1;i++{
		perfix[i] = perfix[i - 1] + candiesCount[i - 1]
	}
	result := make([]bool,len(queries))
	for i:=0;i<len(queries);i++ {
		result[i] = perfix[queries[i][0] + 1] >= queries[i][1] + 1 && perfix[queries[i][0]] < (queries[i][1] + 1) * queries[i][2]
	}
	return result
}


func main() {
	fmt.Println(canEat([]int{46,5,47,48,43,34,15,26,11,25,41,47,15,25,16,50,32,42,32,21,36,34,50,45,46,15,46,38,50,12,3,26,26,16,23,1,4,48,47,32,47,16,33,23,38,2,19,50,6,19,29,3,27,12,6,22,33,28,7,10,12,8,13,24,21,38,43,26,35,18,34,3,14,48,50,34,38,4,50,26,5,35,11,2,35,9,11,31,36,20,21,37,18,34,34,10,21,8,5},[][]int{{11,387,25},{1,3,4},{2,3,3}}))

}
