package ctx

import (
	"testing"
	 "fmt"
)

func TestSetCi(t *testing.T) {

	ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}

	cip := CreateCip().
		SetHeadData(CipHeaderType(HEADER_TYPE_RZV), CIP_ARRAY_RZV).
		SetCi(CiType(CI_TYPE_RZV), CI_BRICK_RZV, ciB).
		SetAppData(AppDataType(APP_DATA_TYPE_RZV), CIP_ARRAY_RZV)


	fmt.Println()
	fmt.Println("type Cip struct { ... }: ")
	fmt.Println(cip)


	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT


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