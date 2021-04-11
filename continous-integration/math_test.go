package main

import "testing"

func TestSum(t *testing.T) {

	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Expect total of the sum to be 30, total = %d", total)
	}
}
