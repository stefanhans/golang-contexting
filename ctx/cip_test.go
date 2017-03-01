package ctx_test

import (
	_ "fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

func TestIt(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"Hello, world", "Hello, world"},
		{"", ""},
	}
	for _, c := range cases {
		got := c.in
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	//ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}
	//
	//cip := CreateCip().
	//	SetHeadData(HEADER_TYPE_RZV, CIP_ARRAY_RZV).
	//	SetCi(CI_TYPE_RZV, CI_BRICK_RZV, ciB).
	//	SetAppData(APP_DATA_TYPE_RZV, CIP_ARRAY_RZV)
	//
	////cip.profile = PROFILE_GATEWAY | PROFILE_ROUTER | PROFILE_REPORTER | PROFILE_STORAGE
	//
	//fmt.Println()
	//fmt.Println("type Cip struct { ... }: ")
	//fmt.Println(cip)
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

	_ = PURPOSE_RZV
	_ = PURPOSE_HEARTBEAT
	_ = PURPOSE_REQUEST
	_ = PURPOSE_OFFER
	_ = PURPOSE_REPLY

	_ = PROFILE_RZV
	_ = PROFILE_ROUTER
	_ = PROFILE_STORAGE
	_ = PROFILE_REPORTER

	_ = CHANNEL_RZV
	_ = CHANNEL_META
	_ = CHANNEL_CONTENT

	_ = HEADER_TYPE_ERROR

	_ = CIP_CI_RZV

	_ = CI_TYPE_SIMPLE_MATCH

	//fmt.Printf("PROFILE_RZV: %08b\n", PROFILE_RZV)
	//fmt.Printf("PROFILE_GATEWAY: %08b\n", PROFILE_GATEWAY)
	//fmt.Printf("PROFILE_ROUTER: %08b\n", PROFILE_ROUTER)
}
