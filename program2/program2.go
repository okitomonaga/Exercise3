package main

import "fmt"

func getdigit(str string) []string {
	search_id := 0
	start_id := 0
	for i := 0; i < len(str); i++ { //1000の位を切り出すためのfor
		if str[search_id] == 'M' {
			search_id++
		} else {
			break
		}
	}
	str_1000 := str[start_id:search_id] //1000の位を切り出す
	start_id = search_id                //start_idを初期化
	for i := start_id; i < len(str); i++ {
		if str[search_id] == 'C' || str[search_id] == 'D' {
			search_id++
		} else {
			if str[search_id] == 'M' {
				search_id++
			}
			break
		}
	}
	str_100 := str[start_id:search_id] //100の位を切り出す
	start_id = search_id
	for i := start_id; i < len(str); i++ {
		if str[search_id] == 'X' || str[search_id] == 'L' {
			search_id++
		} else {
			if str[search_id] == 'C' {
				search_id++
			}
			break
		}
	}
	str_10 := str[start_id:search_id] //10の位を切り出す
	start_id = search_id
	str_1 := str[start_id:len(str)] //残りの1の位を切り出す

	return []string{str_1000, str_100, str_10, str_1}
}

func getnum1000(str string) int {

	var result int
	switch str {
	case "M":
		result = 1000
	case "MM":
		result = 2000
	case "MMM":
		result = 3000
	case "MMMM":
		result = 4000
	default:
		result = 0
	}

	return result

}

func getnum100(str string) int {

	var result int
	switch str {
	case "C":
		result = 100
	case "CC":
		result = 200
	case "CCC":
		result = 300
	case "CD":
		result = 400
	case "D":
		result = 500
	case "DC":
		result = 600
	case "DCC":
		result = 700
	case "DCCC":
		result = 800
	case "CM":
		result = 900
	default:
		result = 0
	}

	return result

}

func getnum10(str string) int {

	var result int
	switch str {
	case "X":
		result = 10
	case "XX":
		result = 20
	case "XXX":
		result = 30
	case "XL":
		result = 40
	case "L":
		result = 50
	case "LX":
		result = 60
	case "LXX":
		result = 70
	case "LXXX":
		result = 80
	case "XC":
		result = 90
	default:
		result = 0
	}

	return result

}

func getnum1(str string) int {

	var result int
	switch str {
	case "I":
		result = 1
	case "II":
		result = 2
	case "III":
		result = 3
	case "IV":
		result = 4
	case "V":
		result = 5
	case "VI":
		result = 6
	case "VII":
		result = 7
	case "VIII":
		result = 8
	case "IX":
		result = 9
	default:
		result = 0
	}

	return result

}

func generate_num(str string) int {

	digit_slice := make([]string, 4)
	copy(digit_slice, getdigit(str))
	result1000 := getnum1000(digit_slice[0])
	result100 := getnum100(digit_slice[1])
	result10 := getnum10(digit_slice[2])
	result1 := getnum1(digit_slice[3])
	return result1000 + result100 + result10 + result1
}

func main() {
	fmt.Println(generate_num("MCMXC"))
	fmt.Println(generate_num("MMVIII"))
	fmt.Println(generate_num("XCIX"))
	fmt.Println(generate_num("XLVII"))
}
