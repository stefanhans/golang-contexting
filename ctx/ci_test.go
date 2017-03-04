package ctx_test

import (
	"fmt"
	_ "fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var ciTestTable_1 = []struct {
	ciType               CiType
	strCiType            string
	rootBrick            CiBrick
	byteRootBrickContent byte
	byteRootBrickMask    byte
}{
	{CI_TYPE_RZV, "CI_TYPE_RZV", CiBrick{0, 0}, 0, 0},
	{CI_TYPE_SIMPLE_MATCH, "CI_TYPE_SIMPLE_MATCH", CiBrick{1, 0}, 1, 0},
	{CiType(255), "CI_TYPE_UNDEFINED", CiBrick{255, 255}, 255, 255},
}

var ciTestTable_2 = []struct {
	ciBricks               CiBrickSlice
}{
	{CiBrickSlice{CiBrick{2, 0}, CiBrick{1, 0}, CiBrick{255, 255} } },
}

func TestCi(t *testing.T) {

	cip := CreateCip()

	for i, ci := range ciTestTable_1 {
		cip.SetCi(ci.ciType, ci.rootBrick, CIP_CI_RZV)
		ciType, rootBrick, _ := cip.Ci()
		s1 := fmt.Sprintf("%s %d %d", ci.ciType, ci.rootBrick.Content, ci.rootBrick.Mask)
		s2 := fmt.Sprintf("%s %d %d", ciType, rootBrick.Content, rootBrick.Mask)
		if s1 != s2 {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s1, s2)
		}
	}

	for i, ci := range ciTestTable_2 {

		var cib CiBricks = CIP_CI_RZV
		cib[0] = CiBrick{byte(len(ci.ciBricks)), 0}
		for i:=1; i<=len(ci.ciBricks); i++ {
			cib[i] = ci.ciBricks[i-1]
		}
		cip.SetCi(CI_TYPE_RZV, CI_BRICK_RZV, cib)
		_, _, ciBricks := cip.Ci()
		s1 := fmt.Sprintf("%s", ci.ciBricks)
		s2 := fmt.Sprintf("%s", ciBricks)
		if s1 != s2 {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s1, s2)
		}
	}

	//ciBricks := CiBricks{CiBrick{1, 2}}
	//cip.SetCi(CI_TYPE_SIMPLE_MATCH, CiBrick{1, 0}, ciBricks)
	//
	//ciType, rootCic, ciBrickArray := cip.Ci()
	//fmt.Printf("%s\n", ciType)
	//fmt.Printf("%s", rootCic)
	//fmt.Printf("%s", ciBrickArray)
}
