package main

import (
	//"reflect"
	"fmt"
	"testing"
)

func TestGetdigit1(t *testing.T) {
	str := "MCXI"
	result := getdigit(str)
	if result[0] != "M" || result[1] != "C" || result[2] != "X" || result[3] != "I" {
		t.Error("Testgetdigit1 is failed")
	}
}

func TestGetdigit2(t *testing.T) {
	str := "MMDLV"
	result := getdigit(str)
	if result[0] != "MM" || result[1] != "D" || result[2] != "L" || result[3] != "V" {
		t.Error("Testgetdigit2 is failed")
	}
}

func TestGetdigit3(t *testing.T) {
	str := "MCMXCIX"
	result := getdigit(str)
	if result[0] != "M" || result[1] != "CM" || result[2] != "XC" || result[3] != "IX" {
		t.Error("Testgetdigit3 is failed")
	}
}

func TestGetdigit4(t *testing.T) {
	str := "MI"
	result := getdigit(str)
	if result[0] != "M" || result[1] != "" || result[2] != "" || result[3] != "I" {
		t.Error("Testgetdigit4 is failed")
	}
}

func TestGetnum1000_1(t *testing.T) {
	str := "M"
	result := getnum1000(str)
	if result != 1000 {
		t.Error("Testgetnum1000_1 is failed")
	}
}

func TestGetnum1000_2(t *testing.T) {
	str := "MMM"
	result := getnum1000(str)
	if result != 3000 {
		t.Error("Testgetnum1000_2 is failed")
	}
}

func TestGetnum100_1(t *testing.T) {
	str := "C"
	result := getnum100(str)
	if result != 100 {
		t.Error("Testgetnum100_1 is failed")
	}
}

func TestGetnum100_2(t *testing.T) {
	str := "D"
	result := getnum100(str)
	if result != 500 {
		t.Error("Testgetnum100_2 is failed")
	}
}

func TestGetnum10_1(t *testing.T) {
	str := "X"
	result := getnum10(str)
	if result != 10 {
		t.Error("Testgetnum10_1 is failed")
	}
}

func TestGetnum10_2(t *testing.T) {
	str := "L"
	result := getnum10(str)
	if result != 50 {
		t.Error("Testgetnum10_2 is failed")
	}
}

func TestGetnum1_1(t *testing.T) {
	str := "I"
	result := getnum1(str)
	if result != 1 {
		t.Error("Testgetnum1_1 is failed")
	}
}

func TestGetnum1_2(t *testing.T) {
	str := "V"
	result := getnum1(str)
	if result != 5 {
		t.Error("Testgetnum1_2 is failed")
	}
}

func TestGenerateNum1_1(t *testing.T) {
	str := "MCXI"
	result := generate_num(str)
	if result != 1111 {
		t.Error("TestGenerateNum1_1 is failed")
	}
}
func TestGenerateNum1_2(t *testing.T) {
	str := "MMDLV"
	result := generate_num(str)
	if result != 2555 {
		t.Error("TestGenerateNum1_2 is failed")
	}
}
