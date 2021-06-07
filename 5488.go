package main

import "fmt"

func minOperations(n int) int {
	if n%2 == 0 {
		return n*n/4
	} else {
		return (n+1)*(n-1)/4
	}
}


func main() {
	fmt.Println(minOperations(3))
	fmt.Println(minOperations(6))
}
