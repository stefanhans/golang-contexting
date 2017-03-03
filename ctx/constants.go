package ctx

import "fmt"

// This software version resp. of the CIP's sender
const (
	MAJOR_RELEASE = 1
	MINOR_RELEASE = 0
)

// PORT constants determine the port number for listening for certain CIPs
const (
	PORT_TCP_META    = 22365
	PORT_UDP_META    = 22366
	PORT_TCP_CONTENT = 22367
	PORT_UDP_CONTENT = 22368
)

// RZV (Reserved Zero Values) are reserved for developing and testing purposes.
const (
	RZV         byte = 0
	CONTENT_RZV      = RZV
	MASK_RZV         = RZV
)

// -----------------------------------------------------------------------------------------------------------------

// PURPOSE constants determine
const (
	PURPOSE_RZV = CipPurpose(iota)
	PURPOSE_HEARTBEAT
	PURPOSE_OFFER
	PURPOSE_REQUEST
	PURPOSE_REPLY
)

// CipPurpose as type resp. "purpose"'as field, in combination with CipChannel resp. "channel", determine what to do with a CIP
type CipPurpose byte

// Implements Stringer() to show purpose of CIP
func (purpose CipPurpose) String() string {
	if purpose == PURPOSE_RZV {
		return "PURPOSE_RZV"
	}
	if purpose == PURPOSE_HEARTBEAT {
		return "PURPOSE_HEARTBEAT"
	}
	if purpose == PURPOSE_OFFER {
		return "PURPOSE_OFFER"
	}
	if purpose == PURPOSE_REQUEST {
		return "PURPOSE_REQUEST"
	}
	if purpose == PURPOSE_REPLY {
		return "PURPOSE_REPLY"
	}
	return "PURPOSE_UNDEFINED"
}

// -----------------------------------------------------------------------------------------------------------------

// PROFILE constants determine the possible roles of the sender of a CIP as flags
const (
	PROFILE_RZV     = CipProfile(0)
	PROFILE_GATEWAY = CipProfile(1 << iota)
	PROFILE_ROUTER
	PROFILE_STORAGE
	PROFILE_REPORTER
)

// Type to link CIP's field profile with constants
type CipProfile byte

// Implements Stringer() to show profile of CIP
func (profile CipProfile) String() string {
	if profile == PROFILE_RZV {
		return "PROFILE_RZV"
	}
	out := ""
	if profile&PROFILE_GATEWAY > 0 {
		out += "PROFILE_GATEWAY"
	}
	if profile&PROFILE_ROUTER > 0 {
		if out == "" {
			out += "PROFILE_ROUTER"
		} else {
			out += " | PROFILE_ROUTER"
		}
	}
	if profile&PROFILE_STORAGE > 0 {
		if out == "" {
			out += "PROFILE_STORAGE"
		} else {
			out += " | PROFILE_STORAGE"
		}
	}
	if profile&PROFILE_REPORTER > 0 {
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


// -----------------------------------------------------------------------------------------------------------------

// VERSION, i.e. <major number>.<minor number> as byte in <4bit>.<4bit> is the software version of the sender
const (
	VERSION CipVersion = MAJOR_RELEASE<<4 + MINOR_RELEASE
)

// Type to link CIP's field version with constants
type CipVersion byte

// Implements Stringer() to show version of CIP
func (version CipVersion) String() string {
	return fmt.Sprintf("%d.%d", (version&0xF0)>>4, version&0x0F)
}


// -----------------------------------------------------------------------------------------------------------------

// CHANNEL constants determine the main topic of the CIP
const (
	CHANNEL_RZV = CipChannel(iota)
	CHANNEL_META
	CHANNEL_CONTENT
)

// Type to link field with constants
type CipChannel byte

// Implements Stringer() to show channel of CIP
func (channel CipChannel) String() string {
	switch channel {
	case CHANNEL_RZV:
		return "CHANNEL_RZV"
	case CHANNEL_META:
		return "CHANNEL_META"
	case CHANNEL_CONTENT:
		return "CHANNEL_CONTENT"
	default:
		return "CHANNEL_UNDEFINED"
	}
}

// -----------------------------------------------------------------------------------------------------------------

// HEADER_TYPE constants determine the type of the dynamic header data
const (
	HEADER_TYPE_RZV = CipHeaderType(iota)
	HEADER_TYPE_CONTENT
	HEADER_TYPE_ERROR
)

// Type to link field with constants
type CipHeaderType byte

// Implements Stringer() to show header type of CIP
func (headerType CipHeaderType) String() string {
	switch headerType {
	case HEADER_TYPE_RZV:
		return "HEADER_TYPE_RZV"
	case HEADER_TYPE_CONTENT:
		return "HEADER_TYPE_CONTENT"
	case HEADER_TYPE_ERROR:
		return "HEADER_TYPE_ERROR"
	default:
		return "HEADER_TYPE_UNDEFINED"
	}
}

//const (
//	ERROR_CATEGORY_NONE = CipErrorCategory(iota)
//)
//
//enum ErrorCategory { ErrorCategoryNone=0, CipFormatError=1, ErrorCategoryUndefined };
//enum ErrorPriority { ErrorPriorityNone=0, ErrorPriorityDebug=1, ErrorPriorityInfo=2, ErrorPriorityNotice=3, ErrorPriorityCritical=4, ErrorPriorityAlert=5, ErrorPriorityEmergency=6, ErrorPriorityUndefined };
//enum CipFormatErrorEnum { CipFormatErrorNone=0, CipFormatErrorOutOfRange=1, CipFormatErrorInconsistent=2, CipFormatErrorWrongProtocol=3, CipFormatErrorUndefined };
//
//
//// Type to link field with constants
//type CipErrorCategory byte
//
//// Implements Stringer() to show header type of CIP
//func (cipErrorCategory CipErrorCategory) String() string {
//
//	switch cipErrorCategory {
//	case HEADER_TYPE_RZV:
//		return "HEADER_TYPE_RZV"
//	case HEADER_TYPE_ERROR:
//		return "HEADER_TYPE_ERROR"
//	default:
//		return "HEADER_TYPE_UNDEFINED"
//	}
//}

// -----------------------------------------------------------------------------------------------------------------

// CI_TYPE constants determine the type of the Contextinformation (CI)
const (
	CI_TYPE_RZV = CiType(iota)
	CI_TYPE_SIMPLE_MATCH
)

// Type to link field with constants
type CiType byte

// Implements Stringer() to show ci type of CIP
func (ciType CiType) String() string {
	switch ciType {
	case CI_TYPE_RZV:
		return "CI_TYPE_RZV"
	case CI_TYPE_SIMPLE_MATCH:
		return "CI_TYPE_SIMPLE_MATCH"
	default:
		return "CI_TYPE_UNDEFINED"
	}
}

// -----------------------------------------------------------------------------------------------------------------

// APP_DATA_TYPE constants determine the type of application data
const (
	APP_DATA_TYPE_RZV = AppDataType(iota)
)

// Type to link field with constants
type AppDataType byte

// Implements Stringer() to show application data type of CIP
func (appDataType AppDataType) String() string {
	switch appDataType {
	case APP_DATA_TYPE_RZV:
		return "APP_DATA_TYPE_RZV"
	default:
		return "APP_DATA_TYPE_UNDEFINED"
	}
}

// RZV (Reserved Zero Value) variables
var (
	CIP_CI_RZV    = CiBricks{CI_BRICK_RZV}
	CIP_ARRAY_RZV = CipArray{0}
)

// CI_BRICK_RZV represents an empty CIBrick
var CI_BRICK_RZV = CiBrick{CONTENT_RZV, MASK_RZV}
