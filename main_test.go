package main

import (
	"testing"
)

func Test_sum(t *testing.T) {
	expected := 40000000000000220
	sum := sum([]int{40000000000000000, 10, 10, 200})
	if sum != expected {
		t.Fatalf("sum error, actual: %d, expected: %d.", sum, expected)
	}
}

func Test_scanDir(t *testing.T) {

}
