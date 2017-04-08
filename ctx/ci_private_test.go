package ctx

import (
	"fmt"
	"testing"
)

var convertToArrayTestTable = []struct {
	ciBrickSlice CiBrickSlice
	ciBrickArray CiBrickArray
}{
	{CiBrickSlice{}, CIP_CI_RZV},
}

var convertToSliceTestTable = []struct {
	ciBrickArray CiBrickArray
	ciBrickSlice CiBrickSlice
}{
	{CIP_CI_RZV, CiBrickSlice{}},
}

func TestConvert(t *testing.T) {

	for i, ci := range convertToArrayTestTable {
		s1 := fmt.Sprintf("%v", ci.ciBrickSlice.toCiBrickArray())
		s2 := fmt.Sprintf("%v", ci.ciBrickArray)
		if s1 != s2 {
			//t.Errorf("%d: Value != Expected:\n%s%s\n", i, s1, s2)
			t.Errorf("%d: ", i)
		}
	}

	for i, ci := range convertToSliceTestTable {
		s1 := fmt.Sprintf("%v", ci.ciBrickArray)
		s2 := fmt.Sprintf("%v", ci.ciBrickSlice.toCiBrickArray())
		if s1 != s2 {
			//t.Errorf("%d: Value != Expected:\n%s%s\n", i, s1, s2)
			t.Errorf("%d: ", i)
		}
	}
}
