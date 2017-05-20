package ctx

/****************************************** FILE COMMENT ******************************************

Implementing the CIP API finally.

ToDo: Design services, API and source files structure

****************************************** FILE COMMENT ******************************************/

import (
	"fmt"
	"net"
	"time"
)

// Cip is the struct for the Contextinformation Pakets (CIP)
type Cip struct {

	// ci_head
	purpose       CipPurpose
	profile       CipProfile
	version       CipVersion
	channel       CipChannel
	uuid          _UUID
	ipAddress     net.IP
	ipPort        int16
	time          int64
	headDataType  CipHeaderType
	headDataSize  byte
	headDataArray []byte

	// ci
	ciType       CiType
	rootCic      CiBrick
	ciSize       byte
	ciBrickArray CiBrickSlice

	// app_data
	appDataType  AppDataType
	appDataSize  byte
	appDataArray []byte
}

func (cip *Cip) Init() {
	fmt.Println("Init()")
	cip.ipAddress = net.IPv4(127, 0, 0, 1)
	cip.ipPort = int16(22365)
	cip.Timestamp(time.Now().Unix())
}

func (cip *Cip) Timestamp(time int64) {
	cip.time = time
}

func (cip *Cip) String() string {

	return fmt.Sprintf("%-16s: %s\n", "purpose", cip.purpose) +
		fmt.Sprintf("%-16s: %s\n", "profile", cip.profile) +
		fmt.Sprintf("%-16s: %s\n", "version", cip.version) +
		fmt.Sprintf("%-16s: %s\n", "channel", cip.channel) +
		fmt.Sprintf("%-16s: %v\n", "uuid", cip.uuid) +
		fmt.Sprintf("%-16s: %v\n", "ipAddress", cip.ipAddress) +
		fmt.Sprintf("%-16s: %v\n", "ipPort", cip.ipPort) +
		fmt.Sprintf("%-16s: %v\n", "time", cip.time) +
		fmt.Sprintf("%-16s: %s\n", "headDataType", cip.headDataType) +
		fmt.Sprintf("%-16s: %d\n", "headDataSize", cip.headDataSize) +
		fmt.Sprintf("%-16s: %v\n", "headDataArray", cip.headDataArray) +
		fmt.Sprintf("%-16s: %s\n", "ciType", cip.ciType) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Content", cip.rootCic.Content) +
		fmt.Sprintf("%-16s: %08b\n", "rootCic Mask", cip.rootCic.Mask) +
		fmt.Sprintf("%-16s: %d\n", "ciSize", cip.ciSize) +
		fmt.Sprintf("%s", cip.ciBrickArray) +
		fmt.Sprintf("%-16s: %s\n", "appDataType", cip.appDataType) +
		fmt.Sprintf("%-16s: %d\n", "appDataSize", cip.appDataSize) +
		fmt.Sprintf("%-16s: %v\n", "appDataArray", cip.appDataArray)
}

// CreateCip creates the initial creation of CIP with UUID and null values
func CreateCip() *Cip {
	return &Cip{
		version: VERSION,
		uuid:    newV1(),
	}
}

// CiMatch returns the match of two Contextinformation (CI), i.e. true or false
func (cip *Cip) CiMatch(anotherCip *Cip) bool {

	return true
}
