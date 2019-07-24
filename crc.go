package crc64

// https://golang.org/src/math/bits/bits.go

// https://crc64.online/
// https://en.wikipedia.org/wiki/Cyclic_redundancy_check
// https://github.com/snksoft/crc/blob/master/crc.go

/*
func reverseBits64(input uint64) uint64 {
	input = input>>1&0x5555555555555555 | input&(0x5555555555555555)<<1
	input = input>>2&0x3333333333333333 | input&(0x3333333333333333)<<2
	input = input>>4&0xF0F0F0F0F0F0F0F | input&(0xF0F0F0F0F0F0F0F)<<4
	input = input>>8&0xFF00FF00FF00FF | input&(0xFF00FF00FF00FF)<<8
	input = input>>16&0xFFFF0000FFFF | input&(0xFFFF0000FFFF)<<16
	return input>>32 | input<<32
}

func reverseByte(input byte) byte {
	input = input>>1&0x55 | input&0x55<<1
	input = input>>2&0x33 | input&0x33<<2
	return input>>4 | input<<4
}*/

// Checksum returns CRC64 checksum for given parameters
func Checksum(data []byte, poly uint64, init uint64, xorout uint64) uint64 {
	var crc uint64 = init
	var bit uint64

	for _, item := range data {
		for j := byte(0x80); j != 0; j >>= 1 {
			if (item & j) != 0 {
				bit = (crc & 0x8000000000000000) ^ 0x8000000000000000
			} else {
				bit = crc & 0x8000000000000000
			}
			switch bit {
			case 0:
				crc <<= 1
			default:
				crc = (crc << 1) ^ poly
			}
		}
	}

	return crc ^ xorout
}
