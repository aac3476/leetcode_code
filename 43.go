package main

import (
	"fmt"
	"math"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m + n)
	for i:=m - 1; i>=0; i-- {
		for j:=n - 1; j>=0; j-- {
			multi := int(num1[i] - '0') * int(num2[j] - '0')
			pos := i + j + 1
			for multi > 0 {
				ansArr[pos] = ansArr[pos] + multi % 10
				multi = int(math.Floor(float64(multi / 10)))
				pos = pos - 1
			}
		}
	}
	for l:=m+n-1; l>=0; l--{
		if ansArr[l] >= 10 {
			ansArr[l-1] = ansArr[l-1] + int(math.Floor(float64(ansArr[l] / 10)))
			ansArr[l] = ansArr[l] % 10
		}
	}
	res := ""
	insert := false
	for l:=0; l<m+n; l++ {
		if insert {
			res = res + strconv.Itoa(ansArr[l])
		} else {
			if ansArr[l] > 0{
				res = res + strconv.Itoa(ansArr[l])
				insert = true
			}
		}
	}
	return res
}

func main() {
	fmt.Println(multiply("498828660196","840477629533"))
}
