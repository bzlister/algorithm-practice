package main

import "fmt"

var mem [][]string

func maxString(p, q string) string {
	if len(p) >= len(q) {
		return p
	} else {
		return q
	}
}

func palindrome(s string, i, j int) string {
	if i > j {
		return ""
	}
	if i == j {
		mem[i][j] = string(s[i])
		return string(s[i])
	}
	if mem[i][j] == "" {
		if s[i] == s[j] {
			mem[i][j] = string(s[i]) + palindrome(s, i+1, j-1) + string(s[j])
		} else {
			mem[i][j] = maxString(palindrome(s, i+1, j), palindrome(s, i, j-1))
		}
	}
	return mem[i][j]
}

func main() {
	s := "aaaracecaraabzkjghki"
	mem = make([][]string, len(s))
	for i := range mem {
		mem[i] = make([]string, len(s))
		for j := range mem[i] {
			mem[i][j] = ""
		}
	}
	pal := palindrome(s, 0, len(s)-1)
	fmt.Println(len(pal))
	fmt.Println(pal)
	fmt.Println(mem)
}
