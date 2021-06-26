package main

import (
	//"reflect"
	//"fmt"
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
