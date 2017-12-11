package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	//HexDigits is number of hexadecimal digits
	HexDigits      int     = 16
	Epsilon        float64 = 1e-17 //a constant for the second part of the sum
	CountTwoPowers int     = 25    //number of power of twos in the table
)

var (
	TableTwoPowers [CountTwoPowers]float64
)

func main() {
	var (
		series1 float64
		//frac, series1, series2, series3, series4 float64
		//position                                 int
		//hexadecimal                              string
	)

	fillTable()

	series1 = hexpi(1, 1000000)

	fmt.Println(series1)
}

func fillTable() {
	TableTwoPowers[0] = 1.0

	for i := 1; i < CountTwoPowers; i++ {
		TableTwoPowers[i] = 2.0 * TableTwoPowers[i-1]
	}
}

func hexpi(jConstant, position int) float64 {
	var (
		denominator, power, sum, term, pos, j float64
	)

	j = float64(jConstant)
	pos = float64(position)

	for i := 0; i < position; i++ {
		k := float64(i)
		denominator = 8.0*k + j
		power = pos - k
		term = base16pow(power, denominator)
		sum = sum + term/denominator

		//remove the integer piece of the sum
		sum = getDecimal(sum)
	}

	//compute an arbitrary amount past the nth term
	for i := position; i < position+50; i++ {
		k := float64(i)
		denominator = 8.0 * k * j

		term = math.Pow(16.0, (pos - k))

		if term < Epsilon {
			break
		}

		sum = sum + term/denominator
		sum = getDecimal(sum)
	}

	fmt.Println(sum)
	return sum
}

func getDecimal(myNumber float64) float64 {
	wholeNum := int64(myNumber)
	return myNumber - float64(wholeNum)
}

func odd(x int) bool {
	if x%2 != 0 {
		return true
	}
	return false
}

func bigpower(x, n int) int {
	r := 1
	y := x

	for n > 1 {
		if odd(n) {
			r = r * y
		}

		temp := math.Floor(float64(n) / 2.0)

		n = int(temp)
		y = y * y
	}
	r = r * y

	result := big.NewInt(int64(r))
	fmt.Println("Big Int:", result)
	return r
}

func base16pow(pos, mod float64) float64 {
	var (
		i                      int
		power1, power2, result float64
	)
	if mod == 1.0 {
		return 0.0
	}

	for i = 0; i < CountTwoPowers; i++ {
		if TableTwoPowers[i] > pos {
			break
		}
	}

	power1 = pos
	power2 = TableTwoPowers[i-1]
	result = 1.0

	//Binary exponentiation algorithm

	for j := 0; j <= i; j++ {
		if power1 >= power2 {
			result = 16.0 * result
			wholeNum := int(result / mod)
			result = result - float64(wholeNum)*mod
			power1 = power1 - power2
		}

		power2 = power2 / 2

		if power2 >= 1.0 {
			result = result * result
			wholeNum := int(result / mod)
			result = result - float64(wholeNum)*mod
		}
	}
	return result
}
