package main

import "testing"

type testData struct {
	input []int
	sum   int
}

func TestSumOfEven(t *testing.T) {
	var inputs []testData
	inputs = append(inputs, testData{[]int{1, 2}, 2})
	inputs = append(inputs, testData{[]int{1, 2, 3, 4, 5}, 6})

	for _, d := range inputs {
		sum := sumOfEven(d.input)
		if sum != d.sum {
			t.Error("expected", d.sum, "got", sum)
		}
	}
}
