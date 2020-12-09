package my_package

import "testing"

type testpair struct {
	values []int
	sum int
}

var tests = []testpair{
	{ []int{1, 2}, 3},
	{[]int{1, 1}, 2},
	{[]int{5, 5}, 10},
}

func TestAdd(t *testing.T) {
	//var v int
	//v = Add(1, 2)
	//
	//if v != 3 {
	//	t.Error("Expected 3, got ", v)
	//}

	for _, pair := range tests {
		v := Add(pair.values[0], pair.values[1])

		if v != pair.sum {
			t.Error("For", pair.values, "expected", pair.sum, "got", v)
		}
	}
}