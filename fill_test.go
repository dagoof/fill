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

	if err = Fill(&c, "a string"); err == nil {
		t.Error("fill should not convert types")
	}

	if err = Fill(b, "a string"); err == nil {
		t.Error("fill should not operate on non pointers")
	}

	if err = Fill(&b, 5); err == nil {
		t.Error("fill should not set on non-assignable types")
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

func ExampleConvertFill() {
	type S string
	var s S

	fmt.Println(Fill(&s, "hello") == nil)
	fmt.Println(s)
	fmt.Println(ConvertFill(&s, "hello") == nil)
	fmt.Println(s)

	// Output:
	// false
	//
	// true
	// hello
}
