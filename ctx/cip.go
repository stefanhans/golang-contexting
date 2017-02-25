package ctx

import "net"

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


// REQUEST REFACTOR?
// enum Request { RequestRZV=0, RequestHeartbeat=1, RequestOffer=2, RequestRequest=2, RequestReply=3, RequestUndefined };

// Service Constants
const (
	SERVICE_RZV = byte(iota)
	SERVICE_HEARTBEAT
	SERVICE_OFFER
	SERVICE_REQUEST
	SERVICE_TCP_REPLY
	SERVICE_UDP_REPLY
	SERVICE_UNDEFINED
)

// Profile Constants
const (
	PROFILE_RZV = byte(iota)
	PROFILE_UNDEFINED
)

// Channel Constants
const (
	CHANNEL_RZV = byte(iota)
	CHANNEL_META
	CHANNEL_CONTENT
	CHANNEL_UNDEFINED
)

// HeaderType Constants
const (
	HEADER_TYPE_RZV = byte(iota)
	HEADER_TYPE_ERROR
	HEADER_TYPE_UNDEFINED
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
	CI_TYPE_UNDEFINED
)


// AppDataType Constants
const (
	APP_DATA_TYPE_RZV = byte(iota)
	APP_DATA_TYPE_UNDEFINED
)



type CiBrick struct {
	content byte
	mask    byte
}

type CipArray [256]byte


// Reserved Zero Value Contextinformation Brick
var (
	CI_BRICK_RZV = CiBrick{CONTENT_RZV, MASK_RZV}
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

// The encoded Contextinformation, i.e. 0 - 255 CiBricks
type CiBricks []CiBrick

// Initial creation of CIP with UUID and null values
func CreateCip() Cip {
	return Cip{
		uuid: newV1(),
	}
}


// Sets the Header Data part of CIP
func (cip Cip) SetHeadData(headDataType byte, headData ...byte) Cip {

	//if len(headData) > 255 {
	//	return errors.New("Length of []byte > 255"), Cip{}
	//}

	return Cip{
		cip.request,
		cip.profile,
		cip.version,
		cip.channel,
		cip.uuid,
		cip.ipAddress,
		cip.time,
		headDataType,
		byte(len(headData)),
		headData,
		cip.ciType,
		cip.rootCic,
		cip.ciSize,
		cip.ciBrickArray,
		cip.appDataType,
		cip.appDataSize,
		cip.appDataArray,
	}
}


// Sets the Contextinformation part of CIP
func (cip Cip) SetCi(ciType byte, rootCic CiBrick, ciBricks ...CiBrick) Cip {

	//if len(ciBricks) > 255 {
	//	return errors.New("Length of []CiBricks > 255"), Cip{}
	//}

	return Cip{
		cip.request,
		cip.profile,
		cip.version,
		cip.channel,
		cip.uuid,
		cip.ipAddress,
		cip.time,
		cip.headDataType,
		cip.headDataSize,
		cip.headDataArray,
		ciType,
		rootCic,
		byte(len(ciBricks)),
		ciBricks,
		cip.appDataType,
		cip.appDataSize,
		cip.appDataArray,
	}
}

func (cip Cip) SetAppData(appDataType byte, appData ...byte) Cip {

	//if len(appData) > 255 {
	//	return errors.New("Length of []byte > 255"), Cip{}
	//}

	return Cip{
		cip.request,
		cip.profile,
		cip.version,
		cip.channel,
		cip.uuid,
		cip.ipAddress,
		cip.time,
		cip.headDataType,
		cip.headDataSize,
		cip.headDataArray,
		cip.ciType,
		cip.rootCic,
		cip.ciSize,
		cip.ciBrickArray,
		appDataType,
		byte(len(appData)),
		appData,
	}
}


func (c Cip) isValid() bool {
	return true
}

func (cb CiBricks) isValid() bool {
	return true
}

// cip is the struct of CIP i.e. Contextinformation Paket
type Cip struct {

	// ci_head
	request       byte
	profile       byte
	version       byte
	channel       byte
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
	ciBrickArray []CiBrick

	// ci_data
	appDataType  byte
	appDataSize  byte
	appDataArray []byte
}
