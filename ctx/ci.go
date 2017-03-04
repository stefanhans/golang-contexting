package ctx

import "fmt"

// Brick for Contextinformation
type CiBrick struct {
	Content byte
	Mask    byte
}

// CI_BRICK_RZV (Reserved Zero Value) represents an empty CIBrick, e.g. for testing of rootCic.
var CI_BRICK_RZV = CiBrick{CONTENT_RZV, MASK_RZV}

// The encoded Contextinformation, i.e. 0 - 255 CiBricks
type CiBricks [256]CiBrick

// CIP_CI_RZV (Reserved Zero Value) with CiBrick.Content 0 determine a quasi empty array for Ci data
var CIP_CI_RZV = CiBricks{CI_BRICK_RZV}

func (ciBrick CiBrick) String() string {
	return fmt.Sprintf("%-16s: %08b\n", "Content", ciBrick.Content) +
		fmt.Sprintf("%-16s: %08b\n", "Mask", ciBrick.Mask)
}

type CiBrickSlice []CiBrick

func (ciBricks CiBrickSlice) String() string {
	out := ""
	for i := 0; i < len(ciBricks); i++ {
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Content", ciBricks[i].Content)
		out += fmt.Sprintf("%-16s: %-3d: %-16s: %08b\n", "CiBrickSlice", i, "Mask", ciBricks[i].Mask)
	}
	return out
}

// SetCi sets the Contextinformation part of CIP
func (cip *Cip) SetCi(ciType CiType, rootCic CiBrick, ciBricks CiBricks) *Cip {
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

// True, if both contents are equal or unequal bits are disabled by set bits in both masks
func (offer CiBrick) ContextMatch(request CiBrick) bool {
	notEqual := offer.Content ^ request.Content
	if notEqual == 0 {
		return true
	}
	offerRelevant := ^notEqual | offer.Mask
	notOfferRelevant := ^offerRelevant
	if notOfferRelevant != 0 {
		return false
	}
	requestRelevant := ^notEqual | request.Mask
	notRequestRelevant := ^requestRelevant
	if notRequestRelevant != 0 {
		return false
	}
	return true
}
