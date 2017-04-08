package ctx

/****************************************** FILE COMMENT ******************************************

Implementing the application data part of CIP except the primitives.

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|   type (1)    |   size (1)    | ............................  | fix
	| ............................................................  | |
	| .......... additional data up to 255 bytes (size) ..........  | dyn
	| ............................................................  | |
	+---------------------------------------------------------------+

ToDo: Finalize

****************************************** FILE COMMENT ******************************************/

// Datastructure to fill the dynamic CIP parts of application
//
// The first byte is the number of the next used bytes (0-255)
type CipAppArray [256]byte

// CIP_ARRAY_RZV (Reserved Zero Value) with 0 as first byte determine a quasi empty array for application data.
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
