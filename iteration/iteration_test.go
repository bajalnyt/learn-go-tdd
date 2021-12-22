package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

//Benchmarking `go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	//b.N --> the framework will determine what is a "good" value for that to let you have some decent results.
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("b", 5)
	fmt.Println(out)
	// Output: bbbbb
}
