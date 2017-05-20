package ctx_test

import (
	"fmt"
	_ "fmt"
	. "github.com/stefanhans/golang-contexting/ctx"
	"testing"
)

func TestEncoding(t *testing.T) {

	cip := CreateCip()
	cip.Init()

	var cipSent, cipReceived string

	// Standard
	cipSent = fmt.Sprintf("%v", cip)

	binaryData, err := cip.MarshalBinary()

	if err != nil {
		t.Errorf("binaryData: \n%v", err)
	} else {
		if err := cip.UnmarshalBinary(binaryData); err != nil {
			t.Errorf("%v", err)
		}
	}
	cipReceived = fmt.Sprintf("%v", cip)
	if cipSent != cipReceived {
		t.Errorf("Received != Sent:\n%s%s\n", cipReceived, cipSent)
	}

	// time = 0
	cip.Timestamp(0)
	cipSent = fmt.Sprintf("%v", cip)

	binaryData, err = cip.MarshalBinary()

	if err != nil {
		t.Errorf("binaryData: \n%v", err)
	} else {
		if err := cip.UnmarshalBinary(binaryData); err != nil {
			t.Errorf("%v", err)
		}
	}
	cipReceived = fmt.Sprintf("%v", cip)
	if cipSent != cipReceived {
		t.Errorf("Received != Sent:\n%s%s\n", cipReceived, cipSent)
	}

	// ciBrickArray
	cip.SetCi(CI_TYPE_RZV, CI_BRICK_RZV, CiBrickArray{CiBrick{1, 0}, CiBrick{0, 0}})
	cipSent = fmt.Sprintf("%v", cip)

	binaryData, err = cip.MarshalBinary()

	if err != nil {
		t.Errorf("binaryData: \n%v", err)
	} else {
		if err := cip.UnmarshalBinary(binaryData); err != nil {
			t.Errorf("%v", err)
		}
	}
	cipReceived = fmt.Sprintf("%v", cip)
	if cipSent != cipReceived {
		t.Errorf("Received != Sent:\n%s%s\n", cipReceived, cipSent)
	}

	// validate() < size
	binaryData = make([]byte, 20)
	if err := cip.UnmarshalBinary(binaryData); err != nil {
		t.Errorf("%v", err)
	}

	// validate() != size
	binaryData, err = cip.MarshalBinary()
	if err != nil {
		t.Errorf("%v", err)
	}
	binaryData = make([]byte, len(binaryData)-2)

	if err := cip.UnmarshalBinary(binaryData); err != nil {
		t.Errorf("%v", err)
	}
}
