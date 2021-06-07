package main

func findMin(arr []int) (int, int, int) {
	min1, min2, p1 := 999, 999, -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < min1 {
			min2 = min1
			min1 = arr[i]
			p1 = i
		} else if arr[i] < min2 {
			min2 = arr[i]
		}
	}
	return min1, min2, p1
}

func minFallingPathSum(arr [][]int) int {
	min1, min2, p1 := findMin(arr[0])
	height := len(arr)
	for i := 1; i < height; i++ {
		nmin1,nmin2,np1:=999,999,-1
		for j := 0; j < len(arr[i]); j++ {
			if j != p1 {
				arr[i][j] += min1
			} else {
				arr[i][j] += min2
			}
			if arr[i][j] < nmin1 {
				nmin2 = nmin1
				nmin1 = arr[i][j]
				np1 = j
			} else if arr[i][j] < min2 {
				nmin2 = arr[i][j]
			}
		}
		min1, min2, p1 = nmin1,nmin2,np1
	}
	result := arr[height-1][0]
	for i:=0;i<len(arr[height-1]);i++{
		if arr[height-1][i] < result{
			result = arr[height-1][i]
		}
	}
	return result
}

func main() {

}
