package apply

import (
	"fmt"
)

func checkExampleError(err error) {
	if err != nil {
		panic(err)
	}
}
func ExampleSuffixExp() {
	prefix := "a+b-a*((c+d)/e-f)+g"
	suffix := NewSuffixExp(prefix)
	fmt.Println(suffix.String())
	// Output:
	// a b + a c d + e / f - * - g +
}

func ExampleEvalSuffixExp() {
	prefix := "1.5+3.4*2-1+3.2*1"
	suffix := NewSuffixExp(prefix)
	v, err := suffix.EvalSuffixExp()
	checkExampleError(err)
	fmt.Printf("%.1f\n",v)
	// Output:
	// 10.5
}

func ExampleIsMatchedParentheses() {
	check := "{(a + b) * (c + d)} * [e,f] + ([{1,2}]) * {{1,2}} + [[(1,2)]]"
	matched := IsMatchedParentheses(check)
	fmt.Println(matched)
	check = "{({{}}](]"
	matched = IsMatchedParentheses(check)
	fmt.Println(matched)
	// Output:
	// true
	// false	
}