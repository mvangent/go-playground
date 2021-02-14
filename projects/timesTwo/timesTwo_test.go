package main

import "testing"

func Test8(t *testing.T) {
	_, nn := timesTwo(8)
	if nn != 16 {
		t.Error(`timesTwo(8) != 16`)
	}
}
