package main

import "math/rand"

func roundToNearest10th(n int) int {
	if n == 0 {
		return 1
	}

	rsp := (n / 10) * 10
	if rsp <= 0 {
		rsp = rand.Intn(n)
	}
	return rsp
}
