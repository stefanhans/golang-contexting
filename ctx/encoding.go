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

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// It encodes the CIP into a binary form and returns the result.
func (cip *Cip) MarshalBinary() (data []byte, err error) {

	data = append(data, byte(cip.purpose))
	data = append(data, byte(cip.profile))
	data = append(data, byte(cip.version))
	data = append(data, byte(cip.channel))

	// TODO: More concise operation to append array
	for b := range cip.uuid {
		data = append(data, cip.uuid[b])
	}

	cip.ipAddress = net.IPv4(127, 0, 0, 1)
	data = append(data, cip.ipAddress[:4]...)
	//fmt.Println(cip.ipAddress)

	bytes2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes2, uint16(int16(cip.ipPort)))
	data = append(data, bytes2...)

	// TODO: Check if branch of time marshalling
	bytes8 := make([]byte, 8)
	if cip.time == 0 {
		cip.time = time.Now().Unix()
		cip.time = int64(binary.LittleEndian.Uint64(bytes8))
	} else {
		binary.LittleEndian.PutUint64(bytes8, uint64(int64(cip.time)))
	}
	data = append(data, bytes8...)

	data = append(data, byte(cip.headDataType))
	data = append(data, byte(cip.headDataSize))
	data = append(data, cip.headDataArray[:cip.headDataSize]...)

	data = append(data, byte(cip.ciType))
	data = append(data, byte(cip.rootCic.Content))
	data = append(data, byte(cip.rootCic.Mask))
	data = append(data, byte(cip.ciSize))
	for i := byte(0); i < cip.ciSize; i++ {
		data = append(data, byte(cip.ciBrickArray[i].Content))
		data = append(data, byte(cip.ciBrickArray[i].Mask))
	}

	data = append(data, byte(cip.appDataType))
	data = append(data, byte(cip.appDataSize))
	data = append(data, cip.appDataArray[:cip.appDataSize]...)

	return data, nil
}

// validate returns an error if the total size according to the dynamic sizes
func (cip *Cip) validate(data []byte) error {
	size := 36 + // static head size
		4 + // static ci size
		2 // static app size

	if len(data) < size {
		return fmt.Errorf("%v bytes is less than %v, i.e. the size of the static values of CIP\n", len(data), size)
	}

	size += int(cip.headDataSize)
	size += int(cip.ciSize * 2)
	size += int(cip.appDataSize)

	if len(data) != size {
		return fmt.Errorf("%v bytes does not match %v, i.e. the size of the static and the dynamic values of CIP\n", len(data), size)
	}

	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It decodes the binary form and updates CIP accordingly.
func (cip *Cip) UnmarshalBinary(data []byte) error {
	if err := cip.validate(data); err != nil {
		return err
	}

	i, j := 0, 0
	//var slice []byte

	cip.purpose = CipPurpose(data[i])
	i++
	cip.profile = CipProfile(data[i])
	i++
	cip.purpose = CipPurpose(data[i])
	i++
	cip.version = CipVersion(data[i])
	i++
	cip.channel = CipChannel(data[i])
	i++
	j = i+16
	//slice = data[i:j]
	var uuid _UUID
	uuid.unmarshalBinary(data[i:j])
	cip.uuid = uuid
	i += j+1

	j = i+8
	//slice = data[i:j]

	var a, b, c, d, e, f, g, h byte = 1, 2, 3, 4, 5 ,6 ,7 ,8

	var i64 int64
	i64 = int64(a<<4)
	i64 = int64(b<<4)
	i64 = int64(c<<4)
	i64 = int64(d<<4)
	i64 = int64(e<<4)
	i64 = int64(f<<4)
	i64 = int64(g<<4)
	i64 = int64(h<<4)
	fmt.Errorf("int64: %v", i64)

	//cip.time = slice
	//i++

	//
	//// ci_head
	//purpose       CipPurpose
	//profile       CipProfile
	//version       CipVersion
	//channel       CipChannel
	//uuid          _UUID
	//ipAddress     net.IP
	//ipPort        int16
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


	return nil
}
