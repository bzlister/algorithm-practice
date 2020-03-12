package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func min(p, q int) int {
	if p <= q {
		return p
	} else {
		return q
	}
}

func median(A []int, rank int) int {
	subproblems := int(math.Ceil(float64(len(A)) / 5))
	medians := make([]int, subproblems)
	for i := 0; i < subproblems; i++ {
		stop := min(5*i+5, len(A))
		sort.Ints(A[5*i : stop])
		medians[i] = A[(stop+5*i)/2]
	}
	pivot := 0
	if len(medians) < 5 {
		sort.Ints(medians)
		pivot = medians[len(medians)/2]
	} else {
		pivot = median(medians, len(medians)/2)
	}
	less := 0
	for k := 0; k < len(A); k++ {
		if A[k] < pivot {
			less++
		}
	}
	L := make([]int, less)
	R := make([]int, len(A)-less-1)
	l := 0
	r := 0
	for x := 0; x < len(A); x++ {
		if A[x] < pivot {
			L[l] = A[x]
			l++
		} else if A[x] > pivot {
			R[r] = A[x]
			r++
		}
	}
	if rank < less {
		return median(L, rank)
	} else if rank > less {
		return median(R, rank-less-1)
	} else {
		return pivot
	}
}

func contains(x int, A []int) bool {
	for _, a := range A {
		if x == a {
			return true
		}
	}
	return false
}

func generateLists(n int) ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	A := make([]int, n)
	B := make([]int, n)
	for t := 0; t < n; t++ {
		randInt := 0
		for contains(randInt, A) {
			randInt = rand.Intn(100 * n)
		}
		A[t] = randInt
		B[t] = randInt
	}
	return A, B
}

func main() {
	A, B := generateLists(10001)
	start := time.Now().UnixNano()
	medA := median(A, len(A)/2)
	fmt.Println(time.Now().UnixNano() - start)
	start = time.Now().UnixNano()
	sort.Ints(B)
	medB := B[len(B)/2]
	fmt.Println(time.Now().UnixNano() - start)
	fmt.Println(medA, medB)
}
