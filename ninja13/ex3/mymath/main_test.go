package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct{
		data []int
		answer float64
	}

	tests := []test{
		test{[]int{10, 20, 40, 60, 90}, 40},
		test{[]int{2, 4, 6, 8, 10, 12}, 7},
		test{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}
	
	for _, v := range tests {
		res := CenteredAvg(v.data)
		if res != v.answer {
			t.Error("Got", res, "expected", v.answer)
		}
	}

}
 
func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	//Output:
	//6
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{123, 744, 140, 200}
	for i := 0 ; i < b.N ; i++ {
		CenteredAvg(xi)
	}
}
