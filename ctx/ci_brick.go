package ctx

/****************************************** FILE COMMENT ******************************************

Implementing the contextinformation primitive CiBrick.

ToDo: Functions for routing etc.

****************************************** FILE COMMENT ******************************************/

import "fmt"

// Brick for Contextinformation
type CiBrick struct {
	Content byte
	Mask    byte
}

// CI_BRICK_RZV (Reserved Zero Value) represents an empty CIBrick, e.g. for testing of rootCic.
var CI_BRICK_RZV = CiBrick{CONTENT_RZV, MASK_RZV}

func (ciBrick CiBrick) String() string {
	return fmt.Sprintf("%-16s: %08b\n", "Content", ciBrick.Content) +
		fmt.Sprintf("%-16s: %08b\n", "Mask", ciBrick.Mask)
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
