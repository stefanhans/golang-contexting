package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var aCip = CreateCip()
var bCip = CreateCip()

var cipTestTable = []struct {
	cip    *Cip
	strCip string
}{
	{aCip, "purpose         : PURPOSE_RZV\n" +
		"profile         : PROFILE_RZV\n" +
		"version         : 1.0\n" +
		"channel         : CHANNEL_RZV\n" +
		"uuid            : " + aCip.Uuid() + "\n" +
		"ipAddress       : <nil>\n" +
		"time            : 0\n" +
		"headDataType    : HEADER_TYPE_RZV\n" +
		"headDataSize    : 0\n" +
		"headDataArray   : []\n" +
		"ciType          : CI_TYPE_RZV\n" +
		"rootCic Content : 00000000\n" +
		"rootCic Mask    : 00000000\n" +
		"ciSize          : 0\n" +
		"appDataType     : APP_DATA_TYPE_RZV\n" +
		"appDataSize     : 0\n" +
		"appDataArray    : []\n"},
}

var ciMatchTestTable = []struct {
	ci_1_Type   CiType
	ci_1_rootBrick   CiBrick
	ci_1_Bricks CiBrickSlice
	ci_2_Type   CiType
	ci_2_rootBrick   CiBrick
	ci_2_Bricks CiBrickSlice
	match bool
}{
	{CI_TYPE_RZV, CI_BRICK_RZV, CiBrickSlice{},
		CI_TYPE_RZV, CI_BRICK_RZV, CiBrickSlice{}, true},
}

func TestCip(t *testing.T) {

	for i, cip := range cipTestTable {
		s := fmt.Sprintf("%s", cip.cip)
		if s != cip.strCip {
			t.Errorf("%d: Value != Expected:\n%s%s\n", i, s, cip.strCip)
		}
	}

	for i, cis := range ciMatchTestTable {
		aCip.SetCi(cis.ci_1_Type, cis.ci_1_rootBrick, CIP_CI_RZV)
		bCip.SetCi(cis.ci_2_Type, cis.ci_2_rootBrick, CIP_CI_RZV)
		if aCip.CiMatch(bCip) != cis.match {
			t.Errorf("%d: ciMatchTestTable", i)
		}
	}

	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT

	_ = CIP_CI_RZV
}
