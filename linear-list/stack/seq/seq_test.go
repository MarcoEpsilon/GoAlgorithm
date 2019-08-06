package seq

import (
	"testing"
)
func checkError(err error, t *testing.T) {
	switch err {
	case nil:
		return
	default:
		t.Error(err)
	}
}
func TestIsResult(t *testing.T) {
	errMsg := "func TestIsResult's result is unexpected"
	is, err := IsResult([]int{1,2,3,4,5},[]int{3,2,5,4,1})
	checkError(err, t)
	if !is {
		t.Error(errMsg)
	}
}