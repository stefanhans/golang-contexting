package ctx


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


//
const (
	PURPOSE_RZV CipPurpose = iota
	PURPOSE_HEARTBEAT
	PURPOSE_OFFER
	PURPOSE_REQUEST
	PURPOSE_REPLY
)


// PROFILE constants determine the possible roles of the sender of a CIP as flags
const (
	PROFILE_RZV     CipProfile = 0
	PROFILE_GATEWAY            = 1 << iota
	PROFILE_ROUTER
	PROFILE_STORAGE
	PROFILE_REPORTER
)

// CHANNEL constants determine the main topic of the CIP
const (
	CHANNEL_RZV CipChannel = iota
	CHANNEL_META
	CHANNEL_CONTENT
)

// HEADER_TYPE constants determine the type of the dynamic part of the header
const (
	HEADER_TYPE_RZV CipHeaderType = iota
	HEADER_TYPE_ERROR
)
/*
enum ErrorCategory { ErrorCategoryNone=0, CipFormatError=1, ErrorCategoryUndefined };
enum ErrorPriority { ErrorPriorityNone=0, ErrorPriorityDebug=1, ErrorPriorityInfo=2, ErrorPriorityNotice=3, ErrorPriorityCritical=4, ErrorPriorityAlert=5, ErrorPriorityEmergency=6, ErrorPriorityUndefined };
enum CipFormatErrorEnum { CipFormatErrorNone=0, CipFormatErrorOutOfRange=1, CipFormatErrorInconsistent=2, CipFormatErrorWrongProtocol=3, CipFormatErrorUndefined };
*/

// CI_TYPE constants determine the type of the Contextinformation (CI)
const (
	CI_TYPE_RZV CiType = iota
	CI_TYPE_SIMPLE_MATCH
)

// APP_DATA_TYPE constants determine the type of application data
const (
	APP_DATA_TYPE_RZV AppDataType = iota
)






// RZV (Reserved Zero Value) variables
var (
	CI_BRICK_RZV  = CiBrick{CONTENT_RZV, MASK_RZV}
	CIP_CI_RZV    = CiBricks{CI_BRICK_RZV}
	CIP_ARRAY_RZV = CipArray{0}
)

