package main

import "fmt"

//import "strconv"
import "math"

func getdigit(num int) []int {

	var digit_slice []int
	for i := 0; i < 4; i++ {
		pow := int(math.Pow(10, float64(3-i)))
		digit_slice = append(digit_slice, num/pow)
		num = num % pow
	}

	return digit_slice
}

func getstring1000(num int) string {

	var result string
	switch num {
	case 1:
		result = "M"
	case 2:
		result = "MM"
	case 3:
		result = "MMM"
	case 4:
		result = "MMMM"
	default:
		result = ""
	}

	return result

}

func getstring100(num int) string {

	var result string
	switch num {
	case 1:
		result = "C"
	case 2:
		result = "CC"
	case 3:
		result = "CCC"
	case 4:
		result = "CD"
	case 5:
		result = "D"
	case 6:
		result = "DC"
	case 7:
		result = "DCC"
	case 8:
		result = "DCCC"
	case 9:
		result = "CM"
	default:
		result = ""
	}

	return result

}

func getstring10(num int) string {

	var result string
	switch num {
	case 1:
		result = "X"
	case 2:
		result = "XX"
	case 3:
		result = "XXX"
	case 4:
		result = "XL"
	case 5:
		result = "L"
	case 6:
		result = "LX"
	case 7:
		result = "LXX"
	case 8:
		result = "LXXX"
	case 9:
		result = "XC"
	default:
		result = ""
	}

	return result

}

func getstring1(num int) string {

	var result string
	switch num {
	case 1:
		result = "I"
	case 2:
		result = "II"
	case 3:
		result = "III"
	case 4:
		result = "IV"
	case 5:
		result = "V"
	case 6:
		result = "VI"
	case 7:
		result = "VII"
	case 8:
		result = "VIII"
	case 9:
		result = "IX"
	default:
		result = ""
	}

	return result

}

func generate_string(num int) string {
	var result string
	digit_slice := make([]int, 4)
	copy(digit_slice, getdigit(num))

	result += getstring1000(digit_slice[0])
	result += getstring100(digit_slice[1])
	result += getstring10(digit_slice[2])
	result += getstring1(digit_slice[3])
	return result
}

func main() {
	fmt.Println(generate_string(1990))
	fmt.Println(generate_string(2008))
	fmt.Println(generate_string(99))
	fmt.Println(generate_string(47))
}
