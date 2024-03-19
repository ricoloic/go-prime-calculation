# Calculating prime numbers

There are methods to calculating prime numbers efficiently, but we can always start with the naive approach and improve it afterward.

## Definition and Usage

We first need to define it before trying to calculate it.

For more information on prime numbers you can take a look at the [wiki page](https://en.wikipedia.org/wiki/Prime_number) for it.

### What is a prime number

Following is a list of aspects defining it:

- [natural number](https://en.wikipedia.org/wiki/Natural_number)
- Greater than `1`
- Not a product of two smaller natural numbers (other than `1`)

Also, we can note that a [composite number](https://en.wikipedia.org/wiki/Composite_number) are the numbers that bridge the gaps between any two consecutive prime numbers. This means that they are a product of two or more smaller natural numbers.

### What is the use of prime numbers

## Finding an algorithm

Let's make an algorithm to find if a given number (`x`) is prime or not.

### The Naive Way

One way to solve the problem could be to iterate starting from `2` upto, but not including, the number we want (`x`).

```go
func main() {
    number := 29
    for i := 2; i < number; i++ {
        // ...
    }
}
```

Here the number we want to find if is prime or not is `29`, which is contained in our `number` variable.

Then for each iteration, we can check the condition if `number` is divisible by `i` (any whole number before) and if so, that means that the `number` is a product of another number other than `1` or itself. This meaning that it is not a prime number.

```go
func main() {
    number := 29
    for i := 2; i < number; i++ {
        if number%i == 0 {
            return false
        }
    }
    return true
}
```

Also, we know that if none of the divisions gives us whole numbers, we have a prime number.

This concludes our naive algorithm implementation.

#### Example

```css
number: 29;
i: 2;  i < number: true; number%i == 0: false;
i: 3;  i < number: true; number%i == 0: false;
i: 4;  i < number: true; number%i == 0: false;
i: 5;  i < number: true; number%i == 0: false;
i: 6;  i < number: true; number%i == 0: false;
i: 7;  i < number: true; number%i == 0: false;
i: 8;  i < number: true; number%i == 0: false;
i: 9;  i < number: true; number%i == 0: false;
i: 10; i < number: true; number%i == 0: false;
i: 11; i < number: true; number%i == 0: false;
i: 12; i < number: true; number%i == 0: false;
i: 13; i < number: true; number%i == 0: false;
i: 14; i < number: true; number%i == 0: false;
i: 15; i < number: true; number%i == 0: false;
i: 16; i < number: true; number%i == 0: false;
i: 17; i < number: true; number%i == 0: false;
i: 18; i < number: true; number%i == 0: false;
i: 19; i < number: true; number%i == 0: false;
i: 20; i < number: true; number%i == 0: false;
i: 21; i < number: true; number%i == 0: false;
i: 22; i < number: true; number%i == 0: false;
i: 23; i < number: true; number%i == 0: false;
i: 24; i < number: true; number%i == 0: false;
i: 25; i < number: true; number%i == 0: false;
i: 26; i < number: true; number%i == 0: false;
i: 27; i < number: true; number%i == 0: false;
i: 28; i < number: true; number%i == 0: false;
return true;
```

### Skipping Ever Other Iteration

We know that any even number are divisible by `2` which means that there is no need to check if the number is divisible by any even numbers other than `2`.

This means that we could reduce the amount of iterations by half by skipping every other one iteration.

But we still have to validate that the `number` is not divisible by `2` to begin with.

```go
func main() {
    number := 29
	
    if number%2 == 0 {
        return false
    }

    for i := 3; i < number; i += 2 {
        if number%i == 0 {
            return false
        }
    }
    return true
}
```

So here the main changes that we did are:

1. Added the condition to validate that our `number` is not divisible by `2`
2. Updated our loop for `i` to start at `2 + 1` (`3`)
3. Updated our loop for `i` to skip the even numbers (`i += 2`)

With this we should get the same results as with our first implementation, but this time we don't have to check even numbers apart from `2`. Meaning, we do, at most, half of the iterations compared to our first implementation.

#### Example

```css
number: 29;
number%2 == 0: false;
i: 3;  i < number: true; number%i == 0: false;
i: 5;  i < number: true; number%i == 0: false;
i: 7;  i < number: true; number%i == 0: false;
i: 9;  i < number: true; number%i == 0: false;
i: 11; i < number: true; number%i == 0: false;
i: 13; i < number: true; number%i == 0: false;
i: 15; i < number: true; number%i == 0: false;
i: 17; i < number: true; number%i == 0: false;
i: 19; i < number: true; number%i == 0: false;
i: 21; i < number: true; number%i == 0: false;
i: 23; i < number: true; number%i == 0: false;
i: 25; i < number: true; number%i == 0: false;
i: 27; i < number: true; number%i == 0: false;
i: 28; i < number: true; number%i == 0: false;
return true;
```

### Keeping Track of Previous Primes

One other thing we could do is to keep track of the previous prime numbers found and only compare our `number` with those.

Similar to the previous implementation where we are skipping all numbers that are products of `2` we can do the same but with all numbers that are a product of other number, this would essentially mean that we skip the check against any composite numbers.

```go
func main() {
    primes := []int{2}
    number := 29

    for i := 3; i < number; i += 2 {
        isPrime := true

        for _, prime := range primes {
            if prime > number { break }

            if number%prime == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            primes = append(primes, number)
        }
		
		if
    }
}
```