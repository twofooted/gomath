# gomath
Personal project for computing difficult or obscure mathematics problems with Go
----------
# Bailey-Borwein-Plouffe Algorithm
The Bailey-Borwein-Plouffe algorithm is used to calculate hexadecimal digits of pi at position p. This is based on the Bailey-Borwein-Plouffe formula which calculates, to great accuracy, the value of pi itself to a given precision. 

# Proth Primes
Proth Primes are special prime numbers found by taking an odd integer and multiplying by 2^n then adding 1. n is iterated from 1 up to the point where a prime number is found. The program here tries to find a Proth Prime for n <= 1000. Example:
```
$ cd ProthPrimes
$ go build
$ ./ProthPrimes 99
Proth Prime has been found!
Prime --> 199
```
