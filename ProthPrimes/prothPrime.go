package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

func main() {
	//check to make sure user has provided a cmd line arg
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command line argument.")
		return
	}

	//1st argument given from running the program converted to integer from ASCII
	arg1 := os.Args[1]
	k, err := strconv.Atoi(arg1)

	//check to make sure the conversion of ASCII to int was valid
	if err != nil {
		log.Fatal(err)
		return
	}

	//check to make sure int provided is not an even number
	if isEven(k) || k < 1 {
		fmt.Println("Proth Prime algorithm only works for odd, positive numbers!")
		return
	}

	//run the Proth Prime function to check for primes, N = k * 2^n + 1
	checkProthPrime(k)
}

func checkProthPrime(k int) {
	//P = k * 2**n + 1
	n := new(big.Float).SetPrec(200)                       //n will be used for the 2^n calculation
	proth := new(big.Int)                                  //final answer. will be set by *big.Float.Int function
	bigK := new(big.Float).SetPrec(200).SetInt64(int64(k)) //convert the supplied k value to big.Float for computation
	one := new(big.Float).SetPrec(200).SetFloat64(1.0)     //set for the "+ 1" part of the equation

	// Check for Proth Primes for n <= 1000
	for i := 0; i < 1000; i++ {

		n.SetFloat64(math.Exp2(float64(i + 1)))
		n.Mul(n, bigK)
		n.Add(n, one)
		n.Int(proth)
		if proth.ProbablyPrime(2) {
			fmt.Println("Proth Prime has been found!")
			fmt.Println("Prime --> ", proth)
			return
		}
	}
	fmt.Println("No Proth Primes. Largest N value computed was ", n)
	return
}

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
