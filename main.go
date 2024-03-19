package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	end := 1_000_000
	primes := []int{2}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filepath := fmt.Sprintf("%s/primes.csv", cwd)

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(file)

	for number := 3; number < end; number += 2 {
		square := int(math.Floor(math.Sqrt(float64(number))))

		isPrime := true
		for _, prime := range primes {
			if prime > square {
				isPrime = true
				break
			}

			if number%prime == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, number)
			_, err := fmt.Fprintf(w, "%d\n", number)
			if err != nil {
				return
			}
		}
	}

	if err := w.Flush(); err != nil {
		panic(err)
	}

	endTime := time.Now()
	fmt.Printf("%v -> calculating primes\n", endTime.Sub(startTime))
}
