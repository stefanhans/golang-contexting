package ctx_test

import (
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
	"fmt"
)

func TestConstantTypes(t *testing.T) {

	fmt.Printf("PURPOSE_REPLY is of type %T\n", PURPOSE_REPLY)
	fmt.Printf("PROFILE_RZV is of type %T\n", PROFILE_RZV)
	fmt.Printf("PROFILE_REPORTER is of type %T\n", PROFILE_REPORTER)
	fmt.Printf("CHANNEL_CONTENT is of type %T\n", CHANNEL_CONTENT)
	fmt.Printf("HEADER_TYPE_ERROR is of type %T\n", HEADER_TYPE_ERROR)
	fmt.Printf("CI_TYPE_SIMPLE_MATCH is of type %T\n", CI_TYPE_SIMPLE_MATCH)
	fmt.Printf("APP_DATA_TYPE_RZV is of type %T\n", APP_DATA_TYPE_RZV)
	fmt.Printf("HEADER_TYPE_ERROR is of type %T\n", HEADER_TYPE_ERROR)
	fmt.Printf("HEADER_TYPE_ERROR is of type %T\n", HEADER_TYPE_ERROR)
	fmt.Printf("HEADER_TYPE_ERROR is of type %T\n", HEADER_TYPE_ERROR)
	fmt.Printf("HEADER_TYPE_ERROR is of type %T\n", HEADER_TYPE_ERROR)

	//cases := []struct {
	//	in, want string
	//}{
	//	{"Hello, world", "Hello, world"},
	//	{"", ""},
	//}
	//for _, c := range cases {
	//	got := c.in
	//	if got != c.want {
	//		t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
	//	}
	//}

	ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}

	cip := CreateCip().
		SetHeadData(HEADER_TYPE_RZV, CIP_ARRAY_RZV).
		SetCi(CI_TYPE_RZV, CI_BRICK_RZV, ciB).
		SetAppData(APP_DATA_TYPE_RZV, CIP_ARRAY_RZV)
	//
	////cip.profile = PROFILE_GATEWAY | PROFILE_ROUTER | PROFILE_REPORTER | PROFILE_STORAGE
	//
	fmt.Println()
	fmt.Println("type Cip struct { ... }: ")
	fmt.Println(cip)

	////fmt.Printf("cip.profile: %08b\n", cip.profile)
	////fmt.Printf("cip.profile: %d\n", cip.profile)
	//
	//fmt.Printf("PROFILE_GATEWAY is of type %T\n", PROFILE_GATEWAY)
	//fmt.Printf("CI_TYPE_SIMPLE_MATCH is of type %T\n", CI_TYPE_SIMPLE_MATCH)
	//fmt.Printf("PROFILE_RZV is of type %T\n", PROFILE_RZV)

	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT



	_ = CIP_CI_RZV

	//fmt.Printf("PROFILE_RZV: %08b\n", PROFILE_RZV)
	//fmt.Printf("PROFILE_GATEWAY: %08b\n", PROFILE_GATEWAY)
	//fmt.Printf("PROFILE_ROUTER: %08b\n", PROFILE_ROUTER)
}
