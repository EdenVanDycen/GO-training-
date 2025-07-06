package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("Ожидалось 6, получилось %d", result)
	}
}
