package ctx

import (
	"fmt"
	"net"
)

// Reserved Zero Value (RZV)
//
// Nearly all of the not well known representations have the so called Reserved Zero Value. That means zero as value is reserved for developing and testing purposes.
const (
	RZV = byte(0)
)

// Reserved Zero Value Contextinformation Content and Mask
const (
	CONTENT_RZV = RZV
	MASK_RZV    = RZV
)

// Ports Constants
const (
	PORT_TCP_META    = 22365
	PORT_UDP_META    = 22366
	PORT_TCP_CONTENT = 22367
	PORT_UDP_CONTENT = 22368
)


// Purpose Constants
const (
	PURPOSE_RZV = byte(iota)
	PURPOSE_HEARTBEAT
	PURPOSE_OFFER
	PURPOSE_REQUEST
	PURPOSE_REPLY
)

// Profile Constants
const (
	PROFILE_RZV = byte(iota)
)

// Version, i.e. <major number>.<minor numbeer> as byte in <4bit><4bit>.
const (
	MAJOR_NUMBER = 1
	MINOR_NUMBER = 2
	VERSION      = CipVersion(MAJOR_NUMBER<<4 + MINOR_NUMBER)
)

// Channel Constants
const (
	CHANNEL_RZV = byte(iota)
	CHANNEL_META
	CHANNEL_CONTENT
)

// HeaderType Constants
const (
	HEADER_TYPE_RZV = byte(iota)
	HEADER_TYPE_ERROR
)

/*
enum ErrorCategory { ErrorCategoryNone=0, CipFormatError=1, ErrorCategoryUndefined };
enum ErrorPriority { ErrorPriorityNone=0, ErrorPriorityDebug=1, ErrorPriorityInfo=2, ErrorPriorityNotice=3, ErrorPriorityCritical=4, ErrorPriorityAlert=5, ErrorPriorityEmergency=6, ErrorPriorityUndefined };
enum CipFormatErrorEnum { CipFormatErrorNone=0, CipFormatErrorOutOfRange=1, CipFormatErrorInconsistent=2, CipFormatErrorWrongProtocol=3, CipFormatErrorUndefined };
*/

// CiType Constants
const (
	CI_TYPE_RZV = byte(iota)
	CI_TYPE_SIMPLE_MATCH
)

// AppDataType Constants
const (
	APP_DATA_TYPE_RZV = byte(iota)
)

// Brick for Contextinformation
type CiBrick struct {
	content byte
	mask    byte
}

// The encoded Contextinformation, i.e. 0 - 255 CiBricks
type CiBricks [256]CiBrick

// Datastructure to fill the dynamic CIP parts of header and application
//
// The first byte is the number of the next used bytes (0-255)
type CipArray [256]byte

// Reserved Zero Value Contextinformation Brick
var (
	CI_BRICK_RZV  = CiBrick{CONTENT_RZV, MASK_RZV}
	CIP_CI_RZV    = CiBricks{CI_BRICK_RZV}
	CIP_ARRAY_RZV = CipArray{0}
)

// True, if both contents are equal or unequal bits are disabled by set bits in both masks
func (offer CiBrick) ContextMatch(request CiBrick) bool {

	notEqual := offer.content ^ request.content
	if notEqual == 0 {
		return true
	}

	offerRelevant := ^notEqual | offer.mask
	notOfferRelevant := ^offerRelevant
	if notOfferRelevant != 0 {
		return false
	}

	requestRelevant := ^notEqual | request.mask
	notRequestRelevant := ^requestRelevant
	if notRequestRelevant != 0 {
		return false
	}
	return true
}

// Initial creation of CIP with UUID and null values
func CreateCip() *Cip {
	return &Cip{
		version: VERSION,
		uuid:    newV1(),
	}
}

// Sets the Header Data part of CIP
func (cip *Cip) SetHeadData(headDataType byte, headData CipArray) *Cip {

	cip.headDataType = headDataType
	cip.headDataSize = headData[0]
	cip.headDataArray = headData[1 : cip.headDataSize+1]

	return cip
}

// Sets the Contextinformation part of CIP
func (cip *Cip) SetCi(ciType byte, rootCic CiBrick, ciBricks CiBricks) *Cip {

	cip.ciType = ciType
	cip.rootCic = rootCic
	cip.ciSize = ciBricks[0].content
	cip.ciBrickArray = ciBricks[1 : cip.ciSize+1]

	return cip
}

func (cip *Cip) SetAppData(appDataType byte, appData CipArray) *Cip {

	cip.appDataType = appDataType
	cip.appDataSize = appData[0]
	cip.appDataArray = appData[1 : cip.appDataSize+1]

	return cip
}

// cip is the struct of CIP i.e. Contextinformation Paket
type Cip struct {

	// ci_head
	purpose       CipPurpose
	profile       CipProfile
	version       CipVersion
	channel       CiChannel
	uuid          _UUID
	ipAddress     net.Addr
	time          int64
	headDataType  byte
	headDataSize  byte
	headDataArray []byte

	// ci
	ciType       byte
	rootCic      CiBrick
	ciSize       byte
	ciBrickArray CiBrickSlice

	// ci_data
	appDataType  byte
	appDataSize  byte
	appDataArray []byte
}


//PURPOSE_RZV
//PURPOSE_HEARTBEAT
//PURPOSE_OFFER
//PURPOSE_REQUEST
//PURPOSE_REPLY

type CipPurpose byte

func (purpose CipPurpose) String() string {

	if purpose == 0 {
		return "PURPOSE_RZV"
	}
	if purpose == 1 {
		return "PURPOSE_HEARTBEAT"
	}
	if purpose == 2 {
		return "PURPOSE_OFFER"
	}
	if purpose == 3 {
		return "PURPOSE_REQUEST"
	}
	if purpose == 4 {
		return "PURPOSE_REPLY"
	}
	return "PURPOSE_UNDEFINED"
}

type CipProfile byte

func (cip Cip) String() string {

	return fmt.Sprintf("%-16s: %08b\n", "purpose", cip.purpose) +
		fmt.Sprintf("%-16s: %08b\n", "profile", cip.profile) +
		fmt.Sprintf("%-16s: %s\n", "version", cip.version) +
		fmt.Sprintf("%-16s: %s\n", "channel", cip.channel) +
		fmt.Sprintf("%-16s: %v\n", "uuid", cip.uuid) +
		fmt.Sprintf("%-16s: %v\n", "ipAddress", cip.ipAddress) +
		fmt.Sprintf("%-16s: %v\n", "time", cip.time) +
		fmt.Sprintf("%-16s: %08b\n", "headDataType", cip.headDataType) +
		fmt.Sprintf("%-16s: %08b\n", "headDataSize", cip.headDataSize) +
		fmt.Sprintf("%-16s: %v\n", "headDataArray", cip.headDataArray) +
		fmt.Sprintf("%-16s: %08b\n", "ciType", cip.ciType) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Content", cip.rootCic.content) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Mask", cip.rootCic.mask) +
		fmt.Sprintf("%-16s: %d\n", "ciSize", cip.ciSize) +

		fmt.Sprintf("%-16s: %v\n", "ciBrickArray", cip.ciBrickArray) +

		fmt.Sprintf("%-16s: %08b\n", "appDataType", cip.appDataType) +
		fmt.Sprintf("%-16s: %d\n", "appDataSize", cip.appDataSize) +
		fmt.Sprintf("%-16s: %v\n", "appDataArray", cip.appDataArray)
}

func (ciBrick CiBrick) String() string {

	return fmt.Sprintf("%-16s: %08b\n", "Content", ciBrick.content) +
		fmt.Sprintf("%-16s: %08b\n", "Mask", ciBrick.mask)
}

type CiBrickSlice []CiBrick

// TODO Debug func (ciBricks CiBricks) String() string (related to "Cip String()")
func (ciBricks CiBrickSlice) String() string {

	out := ""
	for i := 0; i < len(ciBricks); i++ {
		out += fmt.Sprintf("%-3d: %-16s: %08b\n", i, "Content", ciBricks[i].content)
		out += fmt.Sprintf("%-3d: %-16s: %08b\n", i, "Mask", ciBricks[i].mask)
	}
	return out
}

type CipVersion byte

func (version CipVersion) String() string {

	return fmt.Sprintf("%d.%d", (version&^0x0F)>>4, version&^0xF0)
}

//CHANNEL_RZV = byte(iota)
//CHANNEL_META
//CHANNEL_CONTENT
//CHANNEL_UNDEFINED

type CiChannel byte

func (channel CiChannel) String() string {

	if channel == 0 {
		return "CHANNEL_RZV"
	}
	if channel == 1 {
		return "CHANNEL_META"
	}
	if channel == 2 {
		return "CHANNEL_CONTENT"
	}
	return "CHANNEL_UNDEFINED"
}


	//SERVICE_RZV = byte(iota)
	//SERVICE_HEARTBEAT
	//SERVICE_OFFER
	//SERVICE_REQUEST
	//SERVICE_TCP_REPLY
	//SERVICE_UDP_REPLY
	//SERVICE_UNDEFINED


	//PROFILE_RZV = byte(iota)
	//PROFILE_UNDEFINED

