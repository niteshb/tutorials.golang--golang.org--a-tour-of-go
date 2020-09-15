package main

import (
	"fmt"
)

func sqrt(x float64) float64 {
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		diff := z*z - x
		if diff < 0 {
			diff = -diff
		}
		if diff < 0.00000000000001 {
			break
		}
	}
	return z
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%2v: %.10v\n", i, sqrt(float64(i)))
	}
}
