package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	hexpi(2)

	var r = big.NewRat(560, 5465161321)

	fmt.Println(r)
}

func hexpi(pos int) {
	d := float64(pos)
	var result float64

	for i := 0; i < pos; i++ {
		k := float64(i)
		numerator := math.Pow(16, d-k)
		den := 8.0*k + 1
		modulus := math.Mod(numerator, den)
		result = result + (modulus / den)

		if result > 1 {
			result = getDecimal(result)
		}
	}

	something := 37.55643

	c := getDecimal(something)

	fmt.Println(c, result)
}

func getDecimal(myNumber float64) float64 {
	wholeNum := int64(myNumber)
	return myNumber - float64(wholeNum)
}
