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
		//if result != []int{1, 4, 6, 7} {
		t.Error("Testgetdigit1 is failed")
	}
}

func TestGetdigit2(t *testing.T) {
	num := 2456
	result := getdigit(num)
	fmt.Println(result)
	if result[0] != 2 || result[1] != 4 || result[2] != 5 || result[3] != 6 {
		//if result != []int{1, 4, 6, 7} {
		t.Error("Testgetdigit2 is failed")
	}
}
