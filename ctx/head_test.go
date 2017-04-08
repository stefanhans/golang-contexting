package ctx_test

import (
	"fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

var purposeHeadTestTable = []struct {
	purpose    CipPurpose
	strPurpose string
}{
	{PURPOSE_RZV, "PURPOSE_RZV"},
	{PURPOSE_HEARTBEAT, "PURPOSE_HEARTBEAT"},
	{PURPOSE_OFFER, "PURPOSE_OFFER"},
	{PURPOSE_REQUEST, "PURPOSE_REQUEST"},
	{PURPOSE_REPLY, "PURPOSE_REPLY"},
	{CipPurpose(255), "PURPOSE_UNDEFINED"},
}

var profileHeadTestTable = []struct {
	profile    CipProfile
	strProfile string
}{
	{PROFILE_RZV, "PROFILE_RZV"},
	{PROFILE_GATEWAY, "PROFILE_GATEWAY"},
	{PROFILE_ROUTER, "PROFILE_ROUTER"},
	{PROFILE_STORAGE, "PROFILE_STORAGE"},
	{PROFILE_REPORTER, "PROFILE_REPORTER"},
	{CipProfile(255), "PROFILE_GATEWAY | PROFILE_ROUTER | PROFILE_STORAGE | PROFILE_REPORTER"},
	{CipProfile(128), "PROFILE_UNDEFINED"},
}

var channelHeadTestTable = []struct {
	channel    CipChannel
	strChannel string
}{
	{CHANNEL_RZV, "CHANNEL_RZV"},
	{CHANNEL_CONTENT, "CHANNEL_CONTENT"},
	{CHANNEL_META, "CHANNEL_META"},
	{CipChannel(255), "CHANNEL_UNDEFINED"},
}

var headDataTestTable_1 = []struct {
	headDataType    CipHeaderType
	strHeadDataType string
}{
	{HEADER_TYPE_RZV, "HEADER_TYPE_RZV"},
	{HEADER_TYPE_CONTENT, "HEADER_TYPE_CONTENT"},
	{HEADER_TYPE_ERROR, "HEADER_TYPE_ERROR"},
	{CipHeaderType(255), "HEADER_TYPE_UNDEFINED"},
}

var headDataTestTable_2 = []struct {
	headDataArray []byte
}{
	{[]byte{0, 1, 2, 3}},
}

func TestHead(t *testing.T) {

	cip := CreateCip()

	for i, head := range purposeHeadTestTable {
		cip.SetPurpose(head.purpose)
		purpose := cip.Purpose()
		s := fmt.Sprintf("%s", purpose)
		if s != head.strPurpose {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s, head.strPurpose)
		}
	}

	for i, head := range profileHeadTestTable {
		cip.SetProfile(head.profile)
		profile := cip.Profile()
		s := fmt.Sprintf("%s", profile)
		if s != head.strProfile {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s, head.strProfile)
		}
	}

	for i, head := range channelHeadTestTable {
		cip.SetChannel(head.channel)
		channel := cip.Channel()
		s := fmt.Sprintf("%s", channel)
		if s != head.strChannel {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s, head.strChannel)
		}
	}

	for i, head := range headDataTestTable_1 {
		cip.SetHeadData(head.headDataType, CIP_HEAD_ARRAY_RZV)
		headDataType, _ := cip.HeadData()
		s := fmt.Sprintf("%s", headDataType)
		if s != head.strHeadDataType {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s, head.strHeadDataType)
		}
	}

	for i, head := range headDataTestTable_2 {

		var hda CipHeadArray = CIP_HEAD_ARRAY_RZV
		hda[0] = byte(len(head.headDataArray))
		for i := 1; i <= len(head.headDataArray); i++ {
			hda[i] = head.headDataArray[i-1]
		}
		cip.SetHeadData(HEADER_TYPE_RZV, hda)
		_, headDataArray := cip.HeadData()
		s1 := fmt.Sprintf("%s", head.headDataArray)
		s2 := fmt.Sprintf("%s", headDataArray)
		if s1 != s2 {
			t.Errorf("%d: Value != Expected: %s != %s\n", i, s1, s2)
		}
	}
}
