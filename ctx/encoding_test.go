package ctx_test

import (
	_ "fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

func TestEncoding(t *testing.T) {

	cip := CreateCip()

	binaryData, err := cip.MarshalBinary()

	if err != nil {
		t.Errorf("%v", err)
	} else {
		if err := cip.UnmarshalBinary(binaryData); err != nil {
			t.Errorf("%v", err)
		}
	}

	//for i, ci := range CiBrickTestTable {
	//	s := fmt.Sprintf("%s", ci.ciBrick)
	//	if s != ci.strCiBrick {
	//		t.Errorf("%d: Value != Expected:\n%s%s\n", i, s, ci.strCiBrick)
	//	}
	//}

}
