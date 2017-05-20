package ctx

import (
	"encoding/binary"
	"fmt"
	//"net"
	"time"
	"net"
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
	fmt.Println("\nMarshalBinary()")

	data = append(data, byte(cip.purpose))
	fmt.Printf("purpose 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.profile))
	fmt.Printf("profile 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.version))
	fmt.Printf("version 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.channel))
	fmt.Printf("channel 1: len(data): %v\n", len(data))


	// TODO: More concise operation to append array
	for b := range cip.uuid {
		data = append(data, cip.uuid[b])
	}
	fmt.Printf("uuid 16: len(data): %v\n", len(data))

	//cip.ipAddress = net.IPv4(127, 0, 0, 1)
	data = append(data, cip.ipAddress.To4()...)
	//fmt.Println(cip.ipAddress)
	fmt.Printf("ipAddress 4: len(data): %v\n", len(data))

	bytes2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes2, uint16(int16(cip.ipPort)))
	data = append(data, bytes2...)
	fmt.Printf("ipPort 2: len(data): %v\n", len(data))

	// TODO: Check if branch of time marshalling
	bytes8 := make([]byte, 8)
	if cip.time == 0 {
		cip.time = time.Now().Unix()
		cip.time = int64(binary.LittleEndian.Uint64(bytes8))
	} else {
		binary.LittleEndian.PutUint64(bytes8, uint64(int64(cip.time)))
	}
	data = append(data, bytes8...)
	fmt.Printf("time 8: len(data): %v\n", len(data))

	data = append(data, byte(cip.headDataType))
	fmt.Printf("headDataType 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.headDataSize))
	fmt.Printf("headDataSize 1: len(data): %v\n", len(data))

	data = append(data, cip.headDataArray[:cip.headDataSize]...)
	fmt.Printf("headDataArray 2+%v: len(data): %v\n", cip.headDataSize, len(data))

	data = append(data, byte(cip.ciType))
	fmt.Printf("ciType 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.rootCic.Content))
	data = append(data, byte(cip.rootCic.Mask))
	fmt.Printf("rootCic 2: len(data): %v\n", len(data))

	data = append(data, byte(cip.ciSize))
	fmt.Printf("ciSize 1: len(data): %v\n", len(data))

	for i := byte(0); i < cip.ciSize; i++ {
		data = append(data, byte(cip.ciBrickArray[i].Content))
		data = append(data, byte(cip.ciBrickArray[i].Mask))
	}
	fmt.Printf("ciBrickArray 4+(%v*2): len(data): %v\n", cip.ciSize, len(data))

	data = append(data, byte(cip.appDataType))
	fmt.Printf("appDataType 1: len(data): %v\n", len(data))

	data = append(data, byte(cip.appDataSize))
	fmt.Printf("appDataSize 1: len(data): %v\n", len(data))

	data = append(data, cip.appDataArray[:cip.appDataSize]...)
	fmt.Printf("appDataArray 2+%v: len(data): %v\n", cip.appDataSize, len(data))

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
	fmt.Println("\nUnmarshalBinary()")

	fmt.Printf("len(data): %v\n", len(data))

	if err := cip.validate(data); err != nil {
		return err
	}

	i, j := 0, 0

	fmt.Printf("purpose 1: data[%v:%v]\n", i, i+1)
	cip.purpose = CipPurpose(data[i]) // 0
	i++

	fmt.Printf("profile 1: data[%v:%v]\n", i, i+1)
	cip.profile = CipProfile(data[i]) // 1
	i++

	fmt.Printf("version 1: data[%v:%v]\n", i, i+1)
	cip.version = CipVersion(data[i]) // 2
	i++

	fmt.Printf("channel 1: data[%v:%v]\n", i, i+1)
	cip.channel = CipChannel(data[i]) // 3
	i++

	//uuid          _UUID // 4:20
	j = i+16
	fmt.Printf("uuid 16: data[%v:%v]\n", i, j)
	var uuid _UUID
	uuid.unmarshalBinary(data[i:j])
	cip.uuid = uuid
	i = j

	//ipAddress     net.IP // 22:26
	j = i+4
	fmt.Printf("ipAddress 4: data[%v:%v] %v\n", i, j, data[i:j])
	cip.ipAddress = net.IPv4(data[i], data[i+1], data[i+2], data[i+3])
	i += 4

	//ipPort        int16 // 27:29
	j = i+2
	fmt.Printf("ipPort 2: data[%v:%v]\n", i, j)
	cip.ipPort = int16(binary.LittleEndian.Uint16(data[i:j]))
	i = j

	//time          int64 // 30:38
	j = i+8
	fmt.Printf("time 8: data[%v:%v]\n", i, j)
	cip.time = int64(binary.LittleEndian.Uint64(data[i:j]))
	i = j

	//headDataType  CipHeaderType // 39
	fmt.Printf("headDataType 1: data[%v:%v]\n", i, i+1)
	cip.headDataType = CipHeaderType(data[i])
	i++

	//headDataSize  byte // 40
	fmt.Printf("headDataSize 1: data[%v:%v]\n", i, i+1)
	cip.headDataSize = data[i]
	i++

	//headDataArray []byte
	j = i + int(cip.headDataSize)
	fmt.Printf("headDataArray %v: data[%v:%v]\n", cip.headDataSize, i, j)
	cip.headDataArray = data[i:j]
	i = j


	//ciType       CiType
	fmt.Printf("ciType 1: data[%v:%v]\n", i, i+1)
	cip.ciType = CiType(data[i])
	i++

	//rootCic      CiBrick
	fmt.Printf("rootCic 2: data[%v:%v]\n", i, i+2)
	cip.rootCic.Content = data[i]
	i++
	cip.rootCic.Mask = data[i]
	i++

	//ciSize       byte
	fmt.Printf("ciSize 1: data[%v:%v]\n", i, i+1)
	cip.ciSize = data[i]
	i++

	//ciBrickArray CiBrickSlice
	fmt.Printf("ciBrickArray %v: data[%v:%v]\n", cip.ciSize, i, i+int(cip.ciSize)*2)
	for ciBrickNum := byte(0); ciBrickNum<cip.ciSize*2; ciBrickNum += 2 {
		cip.ciBrickArray[ciBrickNum] = CiBrick{data[i], data[i+1]}
	}


	//appDataType  AppDataType
	fmt.Printf("appDataType 1: data[%v:%v]\n", i, i+1)
	cip.appDataType = AppDataType(data[i])
	i++

	//appDataSize  byte
	fmt.Printf("appDataSize 1: data[%v:%v]\n", i, i+1)
	cip.appDataSize = data[i]
	i++

	//appDataArray []byte
	j = i + int(cip.appDataSize)
	fmt.Printf("appDataArray %v: data[%v:%v]\n", cip.appDataSize, i, j)
	cip.appDataArray = data[i:j]




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
	//// app_data
	//appDataType  AppDataType
	//appDataSize  byte
	//appDataArray []byte


	return nil
}
