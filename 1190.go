package main

import "fmt"

type t1190 struct {
	s            []string
	linkPosition []int
	next         []*t1190
	last         *t1190
}

func toNext(s string) t1190 {
	result := t1190{}
	label := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			next := toNext(s[i+1:])
			result.next = append(result.next, &next)
			next.last = &result
			result.linkPosition = append(result.linkPosition, label)
			i += toEnd(s[i:])
			continue
		}
		if s[i] == ')' {
			return result
		}
		label++
		result.s = append(result.s, string(s[i]))
	}
	return result
}

func toEnd(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		}
		if s[i] == ')' {
			count--
		}
		if count == 0 {
			return i
		}
	}
	return len(s)
}

func getPosition(data t1190, j int) (int, int) {
	for i := 0; i < len(data.linkPosition); i++ {
		if data.linkPosition[i] == j {
			count := 0
			for m := i; m < len(data.linkPosition); m++ {
				if data.linkPosition[m] == j {
					count++
				}
			}
			return i, count
		}
	}
	return -1, 0
}

func getResult(data t1190, reverse bool) string {
	result := ""
	if data.s == nil {
		if reverse {
			for i := len(data.next) - 1; i >= 0; i-- {
				result += getResult(*(data.next[i]), !reverse)
			}
		} else {
			for i := 0; i < len(data.next); i++ {
				result += getResult(*(data.next[i]), !reverse)
			}
		}
		return result
	}

	if reverse {
		for i := len(data.s); i >= 0; i-- {
			if i < len(data.s) {
				result = result + data.s[i]
			}
			index, count := getPosition(data, i)
			for j := count - 1; j >= 0; j-- {
				if data.next != nil && index != -1 {
					result = result + getResult(*data.next[index+j], !reverse)
				}
			}

		}
	} else {
		for i := 0; i < len(data.s)+1; i++ {
			index, count := getPosition(data, i)
			for j := 0; j < count; j++ {
				if data.next != nil && index != -1 {
					result = result + getResult(*data.next[index+j], !reverse)
				}
			}
			if i < len(data.s) {
				result = result + data.s[i]
			}
		}
	}
	return result
}

func reverseParentheses(s string) string {
	result := t1190{}
	label := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			next := toNext(s[i+1:])
			result.next = append(result.next, &next)
			next.last = &result
			result.linkPosition = append(result.linkPosition, label)
			i += toEnd(s[i:])
			continue
		}
		label++
		result.s = append(result.s, string(s[i]))
	}
	return getResult(result, false)
}

func main() {
	i := "(abcd)"
	fmt.Println(reverseParentheses(i))
}
