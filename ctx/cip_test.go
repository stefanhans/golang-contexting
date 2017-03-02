package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var purposeTest = []struct {
	constant     CipPurpose
	constType string
	constValue string
}{
	{PURPOSE_RZV, "ctx.CipPurpose", "PURPOSE_RZV"},
	{PURPOSE_HEARTBEAT, "ctx.CipPurpose", "PURPOSE_HEARTBEAT"},
	{PURPOSE_OFFER, "ctx.CipPurpose", "PURPOSE_OFFER"},
	{PURPOSE_REQUEST, "ctx.CipPurpose", "PURPOSE_REQUEST"},
	{PURPOSE_REPLY, "ctx.CipPurpose", "PURPOSE_REPLY"},
}

var profileTest = []struct {
	constant     CipProfile
	constType string
	constValue string
}{
	{PROFILE_RZV, "ctx.CipProfile", "PROFILE_RZV"},
	{PROFILE_GATEWAY, "ctx.CipProfile", "PROFILE_GATEWAY"},
	{PROFILE_ROUTER, "ctx.CipProfile", "PROFILE_ROUTER"},
	{PROFILE_STORAGE, "ctx.CipProfile", "PROFILE_STORAGE"},
	{PROFILE_REPORTER, "ctx.CipProfile", "PROFILE_REPORTER"},
}

var channelTest = []struct {
	constant     CipChannel
	constType string
	constValue string
}{
	{CHANNEL_RZV, "ctx.CipChannel", "CHANNEL_RZV"},
	{CHANNEL_CONTENT, "ctx.CipChannel", "CHANNEL_CONTENT"},
	{CHANNEL_META, "ctx.CipChannel", "CHANNEL_META"},
}

var headerTypeTest = []struct {
	constant     CipHeaderType
	constType string
	constValue string
}{
	{HEADER_TYPE_RZV, "ctx.CipHeaderType", "HEADER_TYPE_RZV"},
	{HEADER_TYPE_CONTENT, "ctx.CipHeaderType", "HEADER_TYPE_CONTENT"},
	{HEADER_TYPE_ERROR, "ctx.CipHeaderType", "HEADER_TYPE_ERROR"},
}

var ciTypeTest = []struct {
	constant     CiType
	constType string
	constValue string
}{
	{CI_TYPE_RZV, "ctx.CiType", "CI_TYPE_RZV"},
	{CI_TYPE_SIMPLE_MATCH, "ctx.CiType", "CI_TYPE_SIMPLE_MATCH"},
}

var appDataTypeTest = []struct {
	constant     AppDataType
	constType string
	constValue string
}{
	{APP_DATA_TYPE_RZV, "ctx.AppDataType", "APP_DATA_TYPE_RZV"},
}

func TestConstants(t *testing.T) {

	for _, purpose := range purposeTest {
		s := fmt.Sprintf("%T %s", purpose.constant, purpose.constant)
		if s != purpose.constType + " " + purpose.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", purpose.constant, purpose.constant, purpose.constType, purpose.constValue)
		}
	}

	for _, profile := range profileTest {
		s := fmt.Sprintf("%T %s", profile.constant, profile.constant)
		if s != profile.constType + " " + profile.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", profile.constant, profile.constant, profile.constType, profile.constValue)
		}
	}

	for _, channel := range channelTest {
		s := fmt.Sprintf("%T %s", channel.constant, channel.constant)
		if s != channel.constType + " " + channel.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", channel.constant, channel.constant, channel.constType, channel.constValue)
		}
	}

	for _, headerType := range headerTypeTest {
		s := fmt.Sprintf("%T %s", headerType.constant, headerType.constant)
		if s != headerType.constType + " " + headerType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", headerType.constant, headerType.constant, headerType.constType, headerType.constValue)
		}
	}

	for _, ciType := range ciTypeTest {
		s := fmt.Sprintf("%T %s", ciType.constant, ciType.constant)
		if s != ciType.constType + " " + ciType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", ciType.constant, ciType.constant, ciType.constType, ciType.constValue)
		}
	}

	for _, appDataType := range appDataTypeTest {
		s := fmt.Sprintf("%T %s", appDataType.constant, appDataType.constant)
		if s != appDataType.constType + " " + appDataType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", appDataType.constant, appDataType.constant, appDataType.constType, appDataType.constValue)
		}
	}
	fmt.Println("TestConstants SUCCESS")

	ciB := CiBricks{CiBrick{2, 0}, CiBrick{1, 2}, CiBrick{2, 3}}

	cip := CreateCip().
		SetHeadData(HEADER_TYPE_RZV, CIP_ARRAY_RZV).
		SetCi(CI_TYPE_RZV, CI_BRICK_RZV, ciB).
		SetAppData(APP_DATA_TYPE_RZV, CIP_ARRAY_RZV)

	fmt.Println()
	fmt.Println("type Cip struct { ... }: ")
	fmt.Println(cip)

	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT

	_ = CIP_CI_RZV
}
