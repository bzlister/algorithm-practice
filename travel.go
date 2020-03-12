package main

import "fmt"

func main() {
	cost := city(8, []int{1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10}, []int{10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1})
	fmt.Println(cost)
}

func min(p, q int) int {
	if p <= q {
		return p
	} else {
		return q
	}
}

func city(c int, P, Q []int) int {
	NY := make([]int, len(P))
	SF := make([]int, len(Q))
	NY[0] = P[0]
	SF[0] = Q[0]
	for i := 1; i < len(P); i++ {
		NY[i] = min(NY[i-1]+P[i], SF[i-1]+Q[i]+c)
		SF[i] = min(SF[i-1]+Q[i], NY[i-1]+P[i]+c)
	}
	return min(NY[len(NY)-1], SF[len(SF)-1])
}
