package ctx

import (
	"fmt"
	"net"
)

// Ports Constants
const (
	PORT_TCP_META    = 22365
	PORT_UDP_META    = 22366
	PORT_TCP_CONTENT = 22367
	PORT_UDP_CONTENT = 22368
)

// Reserved Zero Values (RZV) are reserved for developing and testing purposes.
const (
	RZV         = byte(0)
	CONTENT_RZV = RZV
	MASK_RZV    = RZV
)

// Constants name purpose of CIP
const (
	PURPOSE_RZV CipPurpose = iota
	PURPOSE_HEARTBEAT
	PURPOSE_OFFER
	PURPOSE_REQUEST
	PURPOSE_REPLY
)

// Type to link field with constants
type CipPurpose byte

// Implements Stringer() to show purpose of CIP
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

// Profile constants
const (
	PROFILE_GATEWAY = 1<<iota
	PROFILE_ROUTER
	PROFILE_STORAGE
	PROFILE_REPORTER
	PROFILE_RZV CipProfile = 0
)

// Type to link CIP's field profile with constants
type CipProfile byte

// Implements Stringer() to show profile of CIP
func (profile CipProfile) String() string {

	if profile == 0 {
		return "PROFILE_RZV"
	}

	out := ""
	if profile & 1 > 0 {
		out += "PROFILE_GATEWAY"
	}
	if profile & 2 > 0 {
		if out == "" {
			out += "PROFILE_ROUTER"
		} else {
			out += " | PROFILE_ROUTER"
		}
	}
	if profile & 4 > 0 {
		if out == "" {
			out += "PROFILE_STORAGE"
		} else {
			out += " | PROFILE_STORAGE"
		}
	}
	if profile & 8 > 0 {
		if out == "" {
			out += "PROFILE_REPORTER"
		} else {
			out += " | PROFILE_REPORTER"
		}
	}

	if out == "" {
		return "PROFILE_UNDEFINED"
	}
	return out
}

// Version, i.e. <major number>.<minor numbeer> as byte in <4bit>.<4bit>.
const (
	MAJOR_NUMBER = 1
	MINOR_NUMBER = 2
	VERSION CipVersion = MAJOR_NUMBER<<4 + MINOR_NUMBER
)

// Type to link CIP's field version with constants
type CipVersion byte

// Implements Stringer() to show version of CIP
func (version CipVersion) String() string {

	return fmt.Sprintf("%d.%d", (version&0xF0)>>4, version&0x0F)
}

// Channel constants
const (
	CHANNEL_RZV CipChannel = iota
	CHANNEL_META
	CHANNEL_CONTENT
)

// Type to link field with constants
type CipChannel byte

// Implements Stringer() to show channel of CIP
func (channel CipChannel) String() string {

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

// HeaderType constants
const (
	HEADER_TYPE_RZV CipHeaderType = iota
	HEADER_TYPE_ERROR
)

// Type to link field with constants
type CipHeaderType byte

// Implements Stringer() to show header type of CIP
func (headerType CipHeaderType) String() string {

	if headerType == 0 {
		return "HEADER_TYPE_RZV"
	}
	if headerType == 1 {
		return "HEADER_TYPE_ERROR"
	}
	return "HEADER_TYPE_UNDEFINED"
}

/*
enum ErrorCategory { ErrorCategoryNone=0, CipFormatError=1, ErrorCategoryUndefined };
enum ErrorPriority { ErrorPriorityNone=0, ErrorPriorityDebug=1, ErrorPriorityInfo=2, ErrorPriorityNotice=3, ErrorPriorityCritical=4, ErrorPriorityAlert=5, ErrorPriorityEmergency=6, ErrorPriorityUndefined };
enum CipFormatErrorEnum { CipFormatErrorNone=0, CipFormatErrorOutOfRange=1, CipFormatErrorInconsistent=2, CipFormatErrorWrongProtocol=3, CipFormatErrorUndefined };
*/

// CiType constants
const (
	CI_TYPE_RZV CiType = iota
	CI_TYPE_SIMPLE_MATCH
)

// Type to link field with constants
type CiType byte

// Implements Stringer() to show ci type of CIP
func (ciType CiType) String() string {

	if ciType == 0 {
		return "CI_TYPE_RZV"
	}
	if ciType == 1 {
		return "CI_TYPE_SIMPLE_MATCH"
	}
	return "CI_TYPE_UNDEFINED"
}

// AppDataType constants
const (
	APP_DATA_TYPE_RZV AppDataType = iota
)

// Type to link field with constants
type AppDataType byte

// Implements Stringer() to show application data type of CIP
func (appDataType AppDataType) String() string {

	if appDataType == 0 {
		return "APP_DATA_TYPE_RZV"
	}
	return "APP_DATA_TYPE_UNDEFINED"
}

// Datastructure to fill the dynamic CIP parts of header and application
//
// The first byte is the number of the next used bytes (0-255)
type CipArray [256]byte

// Reserved Zero Value (RZV) variables
var (
	CI_BRICK_RZV  = CiBrick{CONTENT_RZV, MASK_RZV}
	CIP_CI_RZV    = CiBricks{CI_BRICK_RZV}
	CIP_ARRAY_RZV = CipArray{0}
)

// Brick for Contextinformation
type CiBrick struct {
	content byte
	mask    byte
}

// The encoded Contextinformation, i.e. 0 - 255 CiBricks
type CiBricks [256]CiBrick

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
func (cip *Cip) SetHeadData(headDataType CipHeaderType, headData CipArray) *Cip {

	cip.headDataType = headDataType
	cip.headDataSize = headData[0]
	cip.headDataArray = headData[1 : cip.headDataSize+1]

	return cip
}

// Sets the Contextinformation part of CIP
func (cip *Cip) SetCi(ciType CiType, rootCic CiBrick, ciBricks CiBricks) *Cip {

	cip.ciType = ciType
	cip.rootCic = rootCic
	cip.ciSize = ciBricks[0].content
	cip.ciBrickArray = ciBricks[1 : cip.ciSize+1]

	return cip
}

func (cip *Cip) SetAppData(appDataType AppDataType, appData CipArray) *Cip {

	cip.appDataType = appDataType
	cip.appDataSize = appData[0]
	cip.appDataArray = appData[1 : cip.appDataSize+1]

	return cip
}

// Cip is the struct for the Contextinformation Pakets (CIP)
type Cip struct {

	// ci_head
	purpose       CipPurpose
	profile       CipProfile
	version       CipVersion
	channel       CipChannel
	uuid          _UUID
	ipAddress     net.Addr
	time          int64
	headDataType  CipHeaderType
	headDataSize  byte
	headDataArray []byte

	// ci
	ciType       CiType
	rootCic      CiBrick
	ciSize       byte
	ciBrickArray CiBrickSlice

	// ci_data
	appDataType  AppDataType
	appDataSize  byte
	appDataArray []byte
}

func (cip Cip) String() string {

	return fmt.Sprintf("%-16s: %s\n", "purpose", cip.purpose) +
		fmt.Sprintf("%-16s: %s\n", "profile", cip.profile) +
		fmt.Sprintf("%-16s: %s\n", "version", cip.version) +
		fmt.Sprintf("%-16s: %s\n", "channel", cip.channel) +
		fmt.Sprintf("%-16s: %v\n", "uuid", cip.uuid) +
		fmt.Sprintf("%-16s: %v\n", "ipAddress", cip.ipAddress) +
		fmt.Sprintf("%-16s: %v\n", "time", cip.time) +
		fmt.Sprintf("%-16s: %s\n", "headDataType", cip.headDataType) +
		fmt.Sprintf("%-16s: %d\n", "headDataSize", cip.headDataSize) +
		fmt.Sprintf("%-16s: %v\n", "headDataArray", cip.headDataArray) +
		fmt.Sprintf("%-16s: %s\n", "ciType", cip.ciType) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Content", cip.rootCic.content) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Mask", cip.rootCic.mask) +
		fmt.Sprintf("%-16s: %d\n", "ciSize", cip.ciSize) +
		fmt.Sprintf("%s", cip.ciBrickArray) +
		fmt.Sprintf("%-16s: %s\n", "appDataType", cip.appDataType) +
		fmt.Sprintf("%-16s: %d\n", "appDataSize", cip.appDataSize) +
		fmt.Sprintf("%-16s: %v\n", "appDataArray", cip.appDataArray)
}

func (ciBrick CiBrick) String() string {

	return fmt.Sprintf("%-16s: %08b\n", "Content", ciBrick.content) +
		fmt.Sprintf("%-16s: %08b\n", "Mask", ciBrick.mask)
}

type CiBrickSlice []CiBrick

func (ciBricks CiBrickSlice) String() string {

	out := ""
	for i := 0; i < len(ciBricks); i++ {
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Content", ciBricks[i].content)
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Mask", ciBricks[i].mask)
	}
	return out
}
