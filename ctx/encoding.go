package ctx

import (
	"fmt"
	"time"
)

//// ci_head
//purpose       CipPurpose
//profile       CipProfile
//version       CipVersion
//channel       CipChannel
//uuid          _UUID
//ipAddress     net.Addr
//time          int64
//headDataType  CipHeaderType
//headDataSize  byte
//headDataArray []byte
//
//// ci
//ciType       CiType
//rootCic      CiBrick
//ciSize       byte
//ciBrickArray CiBrickSlice
//
//// ci_data
//appDataType  AppDataType
//appDataSize  byte
//appDataArray []byte
//

// validate returns an error if the total size according to the dynamic sizes
func (cip *Cip) validate(data []byte) error {
	size := 36 + // static head size
		4 + // static ci size
		2 // static app size

	if len(data) < size {
		return fmt.Errorf("%v bytes is less than %v, i.e. the size of the static values of CIP\n", len(data), size)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// It encodes the CIP into a binary form and returns the result.
func (cip *Cip) MarshalBinary() (data []byte, err error) {
	var out []byte

	out = append(out, byte(cip.purpose))
	out = append(out, byte(cip.profile))
	out = append(out, byte(cip.version))
	out = append(out, byte(cip.channel))
	// TODO: More concise operation to append array
	for b := range cip.uuid {
		out = append(out, cip.uuid[b])
	}
	//out = append(out, byte(cip.ipAddress.Network()))
	//fmt.Println(byte(cip.ipAddress))
	if cip.time == 0 {
		btime, err := time.Now().MarshalBinary()
		if err != nil {
			return nil, err
		}
		for b := range btime {
			out = append(out, btime[b])
		}

		fmt.Printf("%t\n", btime)
	}
	// out = append(out, byte(cip.time))
	out = append(out, byte(cip.headDataType))
	out = append(out, byte(cip.headDataSize))

	out = append(out, byte(cip.ciType))
	out = append(out, byte(cip.rootCic.Content))
	out = append(out, byte(cip.rootCic.Mask))
	out = append(out, byte(cip.ciSize))

	out = append(out, byte(cip.appDataType))
	out = append(out, byte(cip.appDataSize))

	return out, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It decodes the binary form and updates CIP accordingly.
func (cip *Cip) UnmarshalBinary(data []byte) error {
	if err := cip.validate(data); err != nil {
		return err
	}

	return nil
}
