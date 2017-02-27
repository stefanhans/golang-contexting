package ctx

import (
	"testing"
	 "fmt"
)

func TestSetCi(t *testing.T) {

	//cip := CreateCip().
	//	SetHeadData(HEADER_TYPE_RZV, CIP_ARRAY_RZV).
	//	SetCi(CI_TYPE_RZV, CI_BRICK_RZV, CIP_CI_RZV).
	//	SetAppData(APP_DATA_TYPE_RZV, CIP_ARRAY_RZV)

	//fmt.Println(cip)

	ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}

	cip := CreateCip().
		SetHeadData(CipHeaderType(HEADER_TYPE_RZV), CIP_ARRAY_RZV).
		SetCi(CiType(CI_TYPE_RZV), CI_BRICK_RZV, ciB).
		SetAppData(AppDataType(APP_DATA_TYPE_RZV), CIP_ARRAY_RZV)

	//
	//fmt.Println()
	//fmt.Println("type CipChannel byte: ")
	//fmt.Println(cip.channel)


	fmt.Println()
	fmt.Println("type Cip struct { ... }: ")
	fmt.Println(cip)

	//fmt.Println()
	//fmt.Println("type CiBrick struct { content byte, mask byte}: ")
	//fmt.Println(cip.rootCic)
	//
	//fmt.Println()
	//fmt.Println("type CiBrickSlice []CiBrick: ")
	//fmt.Println(cip.ciBrickArray)

	//fmt.Println()
	//fmt.Println("type CiBricks [256]CiBrick: ")
	//fmt.Println(ciB)



	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT

	_ = CIP_CI_RZV


	_ = PURPOSE_RZV
	_ = PURPOSE_HEARTBEAT
	_ = PURPOSE_REQUEST
	_ = PURPOSE_OFFER
	_ = PURPOSE_REPLY

	_ = PROFILE_RZV

	_ = CHANNEL_RZV
	_ = CHANNEL_META
	_ = CHANNEL_CONTENT

	_ = HEADER_TYPE_ERROR

	_ = CI_TYPE_SIMPLE_MATCH
}