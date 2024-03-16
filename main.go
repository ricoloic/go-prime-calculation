package main

import "math"

func isPrime(primes []int, number int) bool {
	if number == 2 || number == 5 {
		return true
	}

	if number%2 == 0 || number%5 == 0 {
		return false
	}

	square := int(math.Floor(math.Sqrt(float64(number))))
	for _, prime := range primes {
		if prime > square {
			return true
		}

		if number%prime == 0 {
			return false
		}
	}

	return true
}

func main() {
	var primes []int
	end := 1_000_000

	for number := 2; number < end; number++ {
		if isPrime(primes, number) {
			primes = append(primes, number)
			println(number)
		}
	}
}
