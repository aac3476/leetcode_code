package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor{
		return image
	}
	m,n:=len(image),len(image[0])
	dx := []int{1,0,0,-1}
	dy := []int{0,1,-1,0}
	var queue [][]int
	queue = append(queue,[]int{sr,sc})
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	for i:=0;i<len(queue);i++{
		for j:=0;j<4;j++{
			mx := queue[i][0] + dx[j]
			my := queue[i][1] + dy[j]
			if  mx >= 0 && mx < m && my >= 0 && my < n && image[mx][my] == oldColor {
				queue = append(queue,[]int{mx,my})
				image[mx][my]=newColor
			}
		}
	}
	return image
}

func main() {
	i := [][]int{{0,0,0},{0,1,0}}
	fmt.Println(floodFill(i,1,1,2))
}
