package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	//HexDigits is number of hexadecimal digits
	HexDigits      int     = 10
	Epsilon        float64 = 1e-17 //a constant for the second part of the sum
	CountTwoPowers int     = 25    //number of power of twos in the table
	Prec                   = 200   //Precision of the *big.Float types
)

var (
	TableTwoPowers [CountTwoPowers]float64
)

func main() {
	var (
		frac, series1, series2, series3, series4 float64
		position                                 int
		//hexadecimal                              string
	)

	//set the variables
	position = 1000000

	fillTable()

	series1 = hexpi(1, position)
	series2 = hexpi(4, position)
	series3 = hexpi(5, position)
	series4 = hexpi(6, position)

	frac = 4*series1 - 2*series2 - series3 - series4
	frac = getDecimal(frac) + 1

	//0.42342979756754

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

func hexpi(jConstant, position int) *big.Float {
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

func base16pow(pos, mod float64) *big.Float {
	var (
		i              int
		power1, power2 float64
		result         *big.Float
	)
	if mod == 1.0 {
		return new(big.Float).SetPrec(Prec).SetFloat64(0.0)
	}

	for i = 0; i < CountTwoPowers; i++ {
		if TableTwoPowers[i] > pos {
			break
		}
	}

	power1 = pos
	power2 = TableTwoPowers[i-1]
	result = new(big.Float).SetPrec(Prec).SetFloat64(1.0)
	sixteen := new(big.Float).SetPrec(Prec).SetFloat64(16.0)

	//Binary exponentiation algorithm

	for j := 0; j <= i; j++ {
		if power1 >= power2 {
			result = result.Mul(sixteen, result) // 16 * result
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
