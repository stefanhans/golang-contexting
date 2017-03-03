package ctx_test

import (
	_ "fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

func TestCip(t *testing.T) {

	//ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}
	//
	//cip := CreateCip().
	//	SetHeadData(HEADER_TYPE_RZV, CIP_ARRAY_RZV).
	//	SetCi(CI_TYPE_RZV, CI_BRICK_RZV, ciB).
	//	SetAppData(APP_DATA_TYPE_RZV, CIP_ARRAY_RZV)

	//fmt.Println()
	//fmt.Println("type Cip struct { ... }: ")
	//fmt.Println(cip)

	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT

	_ = CIP_CI_RZV
}
