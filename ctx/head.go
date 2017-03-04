package ctx

import "fmt"

// SetPurpose sets the purpose of the CIP
func (cip *Cip) SetPurpose(purpose CipPurpose) *Cip {
	cip.purpose = purpose
	return cip
}

// Purpose returns the purpose of the CIP
func (cip *Cip) Purpose() CipPurpose {
	return cip.purpose
}

// SetProfile sets the profile of the CIP
func (cip *Cip) SetProfile(profile CipProfile) *Cip {
	cip.profile = profile
	return cip
}

// Profile returns the purpose of the CIP
func (cip *Cip) Profile() CipProfile {
	return cip.profile
}

// SetChannel sets the channel of the CIP
func (cip *Cip) SetChannel(channel CipChannel) *Cip {
	cip.channel = channel
	return cip
}

// Channel returns the channel of the CIP
func (cip *Cip) Channel() CipChannel {
	return cip.channel
}

// Uuid returns a string representation of CIP's UUID
func (cip *Cip) Uuid() string {
	return fmt.Sprintf("%v", cip.uuid)
}

// Datastructure to fill the dynamic CIP parts of header and application
//
// The first byte is the number of the next used bytes (0-255)
type CipHeadArray [256]byte

// CIP_ARRAY_RZV (Reserved Zero Value) with 0 as first byte determine a quasi empty array for header or application data.
var CIP_HEAD_ARRAY_RZV = CipHeadArray{RZV}

// Helper construct
type CipHeadArraySlice []byte


// SetHeadData sets the dynamic data of CIP's header
func (cip *Cip) SetHeadData(headDataType CipHeaderType, headData CipHeadArray) *Cip {
	cip.headDataType = headDataType
	cip.headDataSize = headData[0]
	cip.headDataArray = headData[1 : cip.headDataSize+1]
	return cip
}

// HeadData returns the dynamic data of CIP's header
func (cip *Cip) HeadData() (CipHeaderType, CipHeadArraySlice) {
	return cip.headDataType, cip.headDataArray
}