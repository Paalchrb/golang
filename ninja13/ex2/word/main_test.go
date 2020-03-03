package word

import (
	"fmt"
	"testing"

	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/02/02-code-finished/quote"
)

func TestCount (t *testing.T) {
	n := Count("One two three four five six")
	if n != 6 {
		t.Error("Got", n , "expected 6")
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got", v, "expected 1")
			}
		case "two":
			if v != 2 {
				t.Error("Got", v, "expected 2")
			}
		case "three":
			if v != 3 {
				t.Error("Got", v, "expected 3")
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("This is a test string with eight words"))
	//Output:
	//8
}

func BenchmarkCount (b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount (b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		UseCount(quote.SunAlso)
	}
}
