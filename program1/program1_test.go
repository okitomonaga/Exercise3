package main

import (
	//"reflect"
	"fmt"
	"testing"
)

func TestGetdigit1(t *testing.T) {
	num := 1467
	result := getdigit(num)
	fmt.Println(result)
	if result[0] != 1 || result[1] != 4 || result[2] != 6 || result[3] != 7 {
		t.Error("Testgetdigit1 is failed")
	}
}

func TestGetdigit2(t *testing.T) {
	num := 2456
	result := getdigit(num)
	fmt.Println(result)
	if result[0] != 2 || result[1] != 4 || result[2] != 5 || result[3] != 6 {
		t.Error("Testgetdigit2 is failed")
	}
}

func TestGetstring1000_1(t *testing.T) {
	num := 1
	result := getstring1000(num)
	fmt.Println(result)
	if result != "M" {
		t.Error("Testgetsritng1000_1 is failed")
	}
}

func TestGetstring1000_2(t *testing.T) {
	num := 2
	result := getstring1000(num)
	fmt.Println(result)
	if result != "MM" {
		t.Error("Testgetsritng1000_2 is failed")
	}
}

func TestGetstring100_1(t *testing.T) {
	num := 5
	result := getstring100(num)
	fmt.Println(result)
	if result != "D" {
		t.Error("Testgetsritng100_1 is failed")
	}
}

func TestGetstring100_2(t *testing.T) {
	num := 4
	result := getstring100(num)
	fmt.Println(result)
	if result != "CD" {
		t.Error("Testgetsritng100_2 is failed")
	}
}

func TestGetstring10_1(t *testing.T) {
	num := 5
	result := getstring10(num)
	fmt.Println(result)
	if result != "L" {
		t.Error("Testgetsritng10_1 is failed")
	}
}

func TestGetstring10_2(t *testing.T) {
	num := 6
	result := getstring10(num)
	fmt.Println(result)
	if result != "LX" {
		t.Error("Testgetsritng10_2 is failed")
	}
}

func TestGetstring10_3(t *testing.T) {
	num := 0
	result := getstring10(num)
	fmt.Println(result)
	if result != "" {
		t.Error("Testgetsritng10_3 is failed")
	}
}

func TestGetstring1_1(t *testing.T) {
	num := 5
	result := getstring1(num)
	fmt.Println(result)
	if result != "V" {
		t.Error("Testgetsritng1_1 is failed")
	}
}

func TestGetstring1_2(t *testing.T) {
	num := 6
	result := getstring1(num)
	fmt.Println(result)
	if result != "VI" {
		t.Error("Testgetsritng1_2 is failed")
	}
}

func TestGeneratestring01(t *testing.T) {
	num := 1234
	result := generate_string(num)
	fmt.Println(result)
	if result != "MCCXXXIV" {
		t.Error("TestGeneratestring01 is failed")
	}
}

func TestGeneratestring02(t *testing.T) {
	num := 2345
	result := generate_string(num)
	fmt.Println(result)
	if result != "MMCCCXLV" {
		t.Error("TestGeneratestring02 is failed")
	}
}

func TestGeneratestring03(t *testing.T) {
	num := 1024
	result := generate_string(num)
	fmt.Println(result)
	if result != "MXXIV" {
		t.Error("TestGeneratestring03 is failed")
	}
}
