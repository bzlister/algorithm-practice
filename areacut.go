package main

import "fmt"

var (
	a   []int
	b   []int
	c   []int
	mem [][]int
)

func max(nums ...int) int {
	maximum := 0
	for _, num := range nums {
		if maximum < num {
			maximum = num
		}
	}
	return maximum
}

func cut(height, width int) int {
	opt := 0
	if mem[height-1][width-1] == 0 {
		h := 1
		for loopH := true; loopH; loopH = (h <= height/2) {
			w := 1
			for loopW := true; loopW; loopW = (w <= width/2) {
				piece := -1
				for i := 0; i < len(a); i++ {
					if (a[i] <= width && b[i] <= height) || (a[i] <= height && b[i] <= width) {
						piece = max(piece, c[i])
					}
				}
				vert := -1
				horz := -1
				if height > 1 {
					vert = cut(h, width) + cut(height-h, width)
				}
				if width > 1 {
					horz = cut(height, w) + cut(height, width-w)
				}
				opt = max(opt, piece, vert, horz)
				fmt.Printf("Optimal price of %d,%d: %d\n", height, width, opt)
				w++
			}
			h++
		}
		mem[height-1][width-1] = opt
	}
	return mem[height-1][width-1]
}

func main() {
	a = []int{2, 2, 3, 3}
	b = []int{1, 2, 2, 3}
	c = []int{6, 10, 13, 500}
	X := 5
	Y := 5
	mem = make([][]int, X)
	for i := range mem {
		mem[i] = make([]int, Y)
	}
	fmt.Println(cut(X, Y))
}
