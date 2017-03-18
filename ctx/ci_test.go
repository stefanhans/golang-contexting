package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var CiBrickTestTable = []struct {
	ciBrick    CiBrick
	strCiBrick string
}{
	{CI_BRICK_RZV, fmt.Sprintf("%-16s: %08b\n%-16s: %08b\n", "Content", byte(0), "Mask", byte(0))},
	{CiBrick{1, 2}, fmt.Sprintf("%-16s: %08b\n%-16s: %08b\n", "Content", byte(1), "Mask", byte(2))},
	{CiBrick{255, 255}, fmt.Sprintf("%-16s: %08b\n%-16s: %08b\n", "Content", byte(255), "Mask", byte(255))},
}

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
	ciBricks CiBrickSlice
}{
	{CiBrickSlice{CiBrick{2, 0}, CiBrick{1, 0}, CiBrick{255, 255}}},
}

var contextMatchTestTable = []struct {
	ciBrick_1 CiBrick
	ciBrick_2 CiBrick
	match     bool
}{
	{CI_BRICK_RZV, CI_BRICK_RZV, true},
	{CiBrick{35, 0}, CiBrick{35, 0}, true},
	{CiBrick{0, 0}, CiBrick{255, 255}, false},
	{CiBrick{0, 255}, CiBrick{255, 0}, false},
	{CiBrick{0, 255}, CiBrick{255, 255}, true},
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

		var cib CiBrickArray = CIP_CI_RZV
		cib[0] = CiBrick{byte(len(ci.ciBricks)), 0}
		for i := 1; i <= len(ci.ciBricks); i++ {
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

	for i, ci := range contextMatchTestTable {

		if ci.ciBrick_1.ContextMatch(ci.ciBrick_2) != ci.match {
			t.Errorf("%d: Value != Expected:\n%s%s\t\t-> %v\n", i, ci.ciBrick_1, ci.ciBrick_2, ci.match)
		}
	}

	for i, ci := range CiBrickTestTable {
		s := fmt.Sprintf("%s", ci.ciBrick)
		if s != ci.strCiBrick {
			t.Errorf("%d: Value != Expected:\n%s%s\n", i, s, ci.strCiBrick)
		}
	}

}
