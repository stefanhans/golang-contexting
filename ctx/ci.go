package ctx

/****************************************** FILE COMMENT ******************************************

Implementing the contextinformation part of CIP except the primitives.

	0                   1                   2                   3
	0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2
	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	|   type (1)    |         root-CIC (2)          |   size (1)    | fix
	| ............................................................  | |
	| .............. additional data up to 510 bytes .............  | dyn
	| .............. i.e. up to 255 CIC-Bricks  ..................  | |
	| ............................................................  | |
	+---------------------------------------------------------------+

ToDo: Functions for routing etc.

****************************************** FILE COMMENT ******************************************/

import "fmt"

// The encoded Contextinformation, i.e. 0 - 255 CiBrickArray
type CiBrickArray [256]CiBrick

// CIP_CI_RZV (Reserved Zero Value) with CiBrick.Content 0 determine a quasi empty array for Ci data
var CIP_CI_RZV = CiBrickArray{CI_BRICK_RZV}

type CiBrickSlice []CiBrick

func (ciBricks CiBrickSlice) String() string {
	out := ""
	for i := 0; i < len(ciBricks); i++ {
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Content", ciBricks[i].Content)
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Mask", ciBricks[i].Mask)
	}
	return out
}

// toCiBrickArray converts from CiBrickSlice (no length data included) to CiBrickArray (length data included)
func (ciBrickSlice CiBrickSlice) toCiBrickArray() CiBrickArray {
	var ciBrickArray CiBrickArray = CIP_CI_RZV

	//fmt.Printf("CIP_CI_RZV:\n%s\n", CIP_CI_RZV)
	ciBrickArray[0] = CiBrick{byte(len(ciBrickSlice)), 0}
	copy(ciBrickArray[1:byte(len(ciBrickSlice))+1], ciBrickSlice[:])

	//fmt.Printf("ciBrickArray:\n%s\n", ciBrickArray)
	return ciBrickArray
}

// toCiBrickSlice converts from CiBrickArray (length data included) to CiBrickSlice (no length data included)
func (ciBrickArray CiBrickArray) toCiBrickSlice() CiBrickSlice {
	var ciBrickSlice CiBrickSlice = ciBrickArray[1 : ciBrickArray[0].Content+1]

	//fmt.Printf("CIP_CI_RZV:\n%s\n", CIP_CI_RZV)
	//ciBrickArray[0] = CiBrick{byte(len(ciBrickSlice)), 0}
	//copy(ciBrickArray[1:byte(len(ciBrickSlice))+1], ciBrickSlice[:])

	//fmt.Printf("ciBrickArray:\n%s\n", ciBrickArray)
	return ciBrickSlice
}

// SetCi sets the Contextinformation part of CIP
func (cip *Cip) SetCi(ciType CiType, rootCic CiBrick, ciBricks CiBrickArray) *Cip {
	cip.ciType = ciType
	cip.rootCic = rootCic
	cip.ciSize = ciBricks[0].Content
	cip.ciBrickArray = ciBricks[1 : cip.ciSize+1]
	return cip
}

// Ci returns the Contextinformation
func (cip *Cip) Ci() (CiType, CiBrick, CiBrickSlice) {
	return cip.ciType, cip.rootCic, cip.ciBrickArray
}
