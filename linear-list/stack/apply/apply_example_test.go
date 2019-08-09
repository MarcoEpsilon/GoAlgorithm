package apply

import (
	"fmt"
)

func ExampleSuffixExp() {
	prefix := "a+b-a*((c+d)/e-f)+g"
	suffix := NewSuffixExp(prefix)
	fmt.Println(suffix.String())
	// Output:
	// ab+acd+e/f-*-g+
}