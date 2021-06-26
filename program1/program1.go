package main

//import "fmt"
//import "strconv"
import "math"

func getdigit(num int) []int {
	// implementation here
	/**
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(i)
	*/
	var digit_slice []int
	for i := 0; i < 4; i++ {
		pow := int(math.Pow(10, float64(3-i)))
		digit_slice = append(digit_slice, num/pow)
		num = num % pow
	}

	return digit_slice
}

/**
func main() {
	// implementation is not yet
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
*/
