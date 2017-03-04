package ctx



// Datastructure to fill the dynamic CIP parts of header and application
//
// The first byte is the number of the next used bytes (0-255)
type CipAppArray [256]byte

// CIP_ARRAY_RZV (Reserved Zero Value) with 0 as first byte determine a quasi empty array for header or application data.
var CIP_APP_ARRAY_RZV = CipAppArray{RZV}

// Helper construct
type CipAppArraySlice []byte

//  SetAppData sets the application data of CIP
func (cip *Cip) SetAppData(appDataType AppDataType, appData CipAppArray) *Cip {
	cip.appDataType = appDataType
	cip.appDataSize = appData[0]
	cip.appDataArray = appData[1 : cip.appDataSize+1]
	return cip
}

// HeadData returns the dynamic data of CIP's header
func (cip *Cip) AppData() (AppDataType, CipAppArraySlice) {
	return cip.appDataType, cip.appDataArray
}
