package main

import (
	"fmt"
	"math/rand"
	"time"
)

var mem []int

func subsequence(A []int) int {
	if len(A) == 1 {
		return 1
	}
	sub := 1
	max := 1
	for i := 1; i < len(A); i++ {
		if mem[i] == 0 {
			sub = subsequence(A[0:i])
			mem[i] = sub
		} else {
			sub = mem[i]
		}
		if A[i-1] < A[len(A)-1] && sub+1 > max {
			max = sub + 1
		}
	}
	return max
}

func subsequenceWrapper(n int) float64 {
	sequence := make([]int, n)
	for a := 0; a < n; a++ {
		sequence[a] = rand.Intn(n)
	}
	times := make([]int64, 10)
	for i := 1; i < 10; i++ {
		start := time.Now().UnixNano()
		for j := 0; j < i; j++ {
			mem = make([]int, len(sequence))
			subsequence(sequence)
		}
		times[i] = time.Now().UnixNano() - start
	}
	xBar := 4.5
	var sum int64 = 0
	for _, t := range times {
		sum += t
	}
	yBar := float64(sum) / 10.0
	var num float64 = 0.0
	var denom float64 = 0.0
	for k := 0; k < 10; k++ {
		x := float64(k) - xBar
		num += x * (float64(times[k]) - yBar)
		denom += x * x
	}
	return num / denom
}

func main() {
	rand.Seed(time.Now().UnixNano())
	times := make([]float64, 10)
	for i := 0; i < 10; i++ {
		times[i] = subsequenceWrapper(1000 * i)
	}
	fmt.Println(times)
}
