package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestMain(t *testing.T) {

	tests := []test{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2}, 3},
		{[]int{1, 3}, 4},
		{[]int{6, 2, 3}, 11},
	}

	for _, v := range tests {
		x := AddSum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer)
		}
	}

}
