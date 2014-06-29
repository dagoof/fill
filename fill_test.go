package fill

import (
	"fmt"
	"testing"
)

func TestFill(t *testing.T) {
	type S string

	var err error
	var a int
	var b string
	var c S

	if err = Fill(&a, 1); err != nil {
		t.Error(err)
	}

	if err = Fill(&b, "a string"); err != nil {
		t.Error(err)
	}

	if err = Fill(&c, "a string"); err != nil {
		t.Error(err)
	}

}

func ExampleFill() {
	type T struct {
		N int
		S string
	}

	var t T

	Fill(&t, T{5, "hello"})
	fmt.Println(t)
	Fill(&t.N, 20)
	fmt.Println(t)

	// Output:
	// {5 hello}
	// {20 hello}
}
