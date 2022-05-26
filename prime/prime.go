package main

import "fmt"

func main() {
	res := []int{3}
	last := 3
	for {
		last += 2
		prime := true
		for _, v := range res {
			if v*v > last {
				break
			}
			if last%v == 0 {
				prime = false
				break
			}
		}
		if prime {
			res = append(res, last)
			if len(res)%100000 == 0 {
				fmt.Println(last)
			}
			if last > 9999999 {
				break
			}
		}
	}
}
