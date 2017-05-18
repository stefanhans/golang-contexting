package ctx

import (
	"encoding/binary"
	"fmt"
	"net"
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

	size += int(cip.headDataSize)
	size += int(cip.ciSize*2)
	size += int(cip.appDataSize)

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

	cip.ipAddress = net.IPv4(127, 0, 0, 1)
	out = append(out, cip.ipAddress[:4]...)
	fmt.Println(cip.ipAddress)

	bytes2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes2, uint16(int16(cip.ipPort)))
	out = append(out, bytes2...)

	bytes8 := make([]byte, 8)
	if cip.time == 0 {
		cip.time = time.Now().Unix()
		cip.time = int64(binary.LittleEndian.Uint64(bytes8))
	} else {
		binary.LittleEndian.PutUint64(bytes8, uint64(int64(cip.time)))
	}
	out = append(out, bytes8...)

	out = append(out, byte(cip.headDataType))
	out = append(out, byte(cip.headDataSize))
	out = append(out, cip.headDataArray[:cip.headDataSize]...)

	out = append(out, byte(cip.ciType))
	out = append(out, byte(cip.rootCic.Content))
	out = append(out, byte(cip.rootCic.Mask))
	out = append(out, byte(cip.ciSize))
	for i:=byte(0); i<cip.ciSize; i++ {
		out = append(out, byte(cip.ciBrickArray[i].Content))
		out = append(out, byte(cip.ciBrickArray[i].Mask))
	}

	out = append(out, byte(cip.appDataType))
	out = append(out, byte(cip.appDataSize))
	out = append(out, cip.appDataArray[:cip.appDataSize]...)

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
