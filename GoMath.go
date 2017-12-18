package main

import (
	"fmt"
	"math"
)

const (
	HexDigits      int     = 10    //HexDigits is number of hexadecimal digits
	Epsilon        float64 = 1e-17 //a constant for the second part of the sum
	CountTwoPowers int     = 25    //number of power of twos in the table
)

var (
	TableTwoPowers [CountTwoPowers]float64 //table of powers of two up to 2**CountTwoPowers
)

func main() {
	var (
		frac, series1, series2, series3, series4 float64
		position                                 int
	)

	//set the variables
	position = 1000000

	//set the power of twos table
	fillTable()

	series1 = hexpi(1, position)
	series2 = hexpi(4, position)
	series3 = hexpi(5, position)
	series4 = hexpi(6, position)

	frac = 4*series1 - 2*series2 - series3 - series4
	frac = getDecimal(frac) + 1

	//0.42342979756754 should be the answer for p=1000000

	hexadecimals := hexString(frac, HexDigits)

	fmt.Println("Hexadecimals:", hexadecimals)
}

func fillTable() {
	TableTwoPowers[0] = 1.0

	for i := 1; i < CountTwoPowers; i++ {
		TableTwoPowers[i] = 2.0 * TableTwoPowers[i-1]
	}
}

func hexString(x float64, digits int) string {
	var (
		reference = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
		result    string
	)

	for i := 0; i < digits; i++ {
		y := 16.0 * x
		floor := int(math.Floor(y))
		result = result + reference[floor]
		fmt.Println(x, y)
		x = getDecimal(y)
	}

	return result
}

func hexpi(jConstant, position int) float64 {
	var (
		denominator, power float64
		sum, term, pos, j  float64
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
	for i := position; i < position+1000; i++ {
		k := float64(i)

		denominator = 8.0 * k * j

		term = math.Pow(16.0, (pos-k)) / denominator

		if term < Epsilon {
			break
		}

		sum = sum + term
		sum = getDecimal(sum)
	}

	return sum
}

func getDecimal(myNumber float64) float64 {
	wholeNum := int(myNumber)
	return myNumber - float64(wholeNum)
}

func base16pow(pos, mod float64) float64 {
	var (
		i              int
		power1, power2 float64
		result         float64
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
			result = 16 * result
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
