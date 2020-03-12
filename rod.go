package main

import (
	"fmt"
)

func Max(p, q int) int {
	if p >= q {
		return p
	} else {
		return q
	}
}

func maxval(n, c int, p []int) int {
	//p[i] = profit for cut of length (i+1)
	val := make([]int, n+1)
	for i := 1; i <= n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			fmt.Printf("%d: Max(%d, %d)\n", i, max, p[j]+val[i-j-1])
			max = Max(max, p[j]+val[i-j-1]-c)
		}
		val[i] = max
	}
	fmt.Println(val)
	return val[n]

}

func main() {
	fmt.Println(maxval(8, 1, []int{1, 5, 8, 9, 10, 17, 17, 20}))
}
