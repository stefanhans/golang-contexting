package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var purposeTestTable = []struct {
	constant   CipPurpose
	constType  string
	constValue string
}{
	{PURPOSE_RZV, "ctx.CipPurpose", "PURPOSE_RZV"},
	{PURPOSE_HEARTBEAT, "ctx.CipPurpose", "PURPOSE_HEARTBEAT"},
	{PURPOSE_OFFER, "ctx.CipPurpose", "PURPOSE_OFFER"},
	{PURPOSE_REQUEST, "ctx.CipPurpose", "PURPOSE_REQUEST"},
	{PURPOSE_REPLY, "ctx.CipPurpose", "PURPOSE_REPLY"},
}

var profileTestTable = []struct {
	constant   CipProfile
	constType  string
	constValue string
}{
	{PROFILE_RZV, "ctx.CipProfile", "PROFILE_RZV"},
	{PROFILE_GATEWAY, "ctx.CipProfile", "PROFILE_GATEWAY"},
	{PROFILE_ROUTER, "ctx.CipProfile", "PROFILE_ROUTER"},
	{PROFILE_STORAGE, "ctx.CipProfile", "PROFILE_STORAGE"},
	{PROFILE_REPORTER, "ctx.CipProfile", "PROFILE_REPORTER"},
	{CipProfile(255), "ctx.CipProfile", "PROFILE_GATEWAY | PROFILE_ROUTER | PROFILE_STORAGE | PROFILE_REPORTER"},
	{CipProfile(128), "ctx.CipProfile", "PROFILE_UNDEFINED"},
}

var versionTestTable = []struct {
	constant        CipVersion
	constType       string
	constValueMajor byte
	constValueMinor byte
}{
	{VERSION, "ctx.CipVersion", MAJOR_RELEASE, MINOR_RELEASE},
}

var channelTestTable = []struct {
	constant   CipChannel
	constType  string
	constValue string
}{
	{CHANNEL_RZV, "ctx.CipChannel", "CHANNEL_RZV"},
	{CHANNEL_CONTENT, "ctx.CipChannel", "CHANNEL_CONTENT"},
	{CHANNEL_META, "ctx.CipChannel", "CHANNEL_META"},
}

var headerTypeTestTable = []struct {
	constant   CipHeaderType
	constType  string
	constValue string
}{
	{HEADER_TYPE_RZV, "ctx.CipHeaderType", "HEADER_TYPE_RZV"},
	{HEADER_TYPE_CONTENT, "ctx.CipHeaderType", "HEADER_TYPE_CONTENT"},
	{HEADER_TYPE_ERROR, "ctx.CipHeaderType", "HEADER_TYPE_ERROR"},
}

var ciTypeTestTable = []struct {
	constant   CiType
	constType  string
	constValue string
}{
	{CI_TYPE_RZV, "ctx.CiType", "CI_TYPE_RZV"},
	{CI_TYPE_SIMPLE_MATCH, "ctx.CiType", "CI_TYPE_SIMPLE_MATCH"},
}

var appDataTypeTestTable = []struct {
	constant   AppDataType
	constType  string
	constValue string
}{
	{APP_DATA_TYPE_RZV, "ctx.AppDataType", "APP_DATA_TYPE_RZV"},
}

func TestConstants(t *testing.T) {

	for _, purpose := range purposeTestTable {
		s := fmt.Sprintf("%T %s", purpose.constant, purpose.constant)
		if s != purpose.constType+" "+purpose.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", purpose.constant, purpose.constant, purpose.constType, purpose.constValue)
		}
	}
	if s := CipPurpose(255).String(); s != "PURPOSE_UNDEFINED" {
		t.Errorf("CipPurpose(255).String(): %s != %s\n", s, "PURPOSE_UNDEFINED")
	}

	for _, profile := range profileTestTable {
		s := fmt.Sprintf("%T %s", profile.constant, profile.constant)
		if s != profile.constType+" "+profile.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", profile.constant, profile.constant, profile.constType, profile.constValue)
		}
	}

	for _, version := range versionTestTable {
		s1 := fmt.Sprintf("%T %s", version.constant, version.constant)
		s2 := fmt.Sprintf("%s %d.%d", version.constType, version.constValueMajor, version.constValueMinor)
		if s1 != s2 {
			t.Errorf("Type Value: %T %s != %s %d.%d\n", version.constant, version.constant, version.constType, version.constValueMajor, version.constValueMinor)
		}
	}

	for _, channel := range channelTestTable {
		s := fmt.Sprintf("%T %s", channel.constant, channel.constant)
		if s != channel.constType+" "+channel.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", channel.constant, channel.constant, channel.constType, channel.constValue)
		}
	}
	if s := CipChannel(255).String(); s != "CHANNEL_UNDEFINED" {
		t.Errorf("CipChannel(255).String(): %s != %s\n", s, "CHANNEL_UNDEFINED")
	}

	for _, headerType := range headerTypeTestTable {
		s := fmt.Sprintf("%T %s", headerType.constant, headerType.constant)
		if s != headerType.constType+" "+headerType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", headerType.constant, headerType.constant, headerType.constType, headerType.constValue)
		}
	}
	if s := CipHeaderType(255).String(); s != "HEADER_TYPE_UNDEFINED" {
		t.Errorf("CipHeaderType(255).String(): %s != %s\n", s, "HEADER_TYPE_UNDEFINED")
	}

	for _, ciType := range ciTypeTestTable {
		s := fmt.Sprintf("%T %s", ciType.constant, ciType.constant)
		if s != ciType.constType+" "+ciType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", ciType.constant, ciType.constant, ciType.constType, ciType.constValue)
		}
	}
	if s := CiType(255).String(); s != "CI_TYPE_UNDEFINED" {
		t.Errorf("CiType(255).String(): %s != %s\n", s, "CI_TYPE_UNDEFINED")
	}

	for _, appDataType := range appDataTypeTestTable {
		s := fmt.Sprintf("%T %s", appDataType.constant, appDataType.constant)
		if s != appDataType.constType+" "+appDataType.constValue {
			t.Errorf("Type Value: %T %s != %s %s\n", appDataType.constant, appDataType.constant, appDataType.constType, appDataType.constValue)
		}
	}
	if s := AppDataType(255).String(); s != "APP_DATA_TYPE_UNDEFINED" {
		t.Errorf("AppDataType(255).String(): %s != %s\n", s, "APP_DATA_TYPE_UNDEFINED")
	}

	_ = PORT_TCP_META
	_ = PORT_UDP_META
	_ = PORT_TCP_CONTENT
	_ = PORT_UDP_CONTENT

	_ = CIP_CI_RZV
}
