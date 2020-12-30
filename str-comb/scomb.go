//////////////////////////////
// scomb.go
package main

import (
	"fmt"
	"strconv"
)

func newGap(gap int) int {
	gap = int(float64(gap) / 1.3)
	if gap == 9 || gap == 10 {
		return 11
	}
	if gap < 1 {
		return 1
	}
	return gap
}
func combSort(a []string) {
	xlen := len(a)
	gap := xlen
	swapped := false
	for {
		swapped = false
		gap = newGap(gap)
		for i := 0; i < xlen-gap; i++ {
			if a[i] > a[i+gap] {
				swapped = true
				a[i], a[i+gap] = a[i+gap], a[i]
			}
		}
		if !(gap > 1 || swapped) {
			break
		}
	}
}

const N = 10000000

func main() {
	arr := make([]string, N, N)
	for z := 0; z < N; z++ {
		arr[z] = strconv.Itoa(N - z)
	}
	combSort(arr)
	for z := 1; z < N; z++ {
		if arr[z] < arr[z-1] {
			fmt.Print("!")
		}
	}
}
