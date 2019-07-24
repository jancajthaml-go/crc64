package crc64

import (
	"strings"
	"testing"
)

var largeText = []byte(strings.Repeat("a", 50000))
var smallText = []byte(strings.Repeat("a", 5))

func AssetEqual(t *testing.T, expected uint64, actual uint64) {
	if expected != actual {
		t.Errorf("Expected 0x%016X got 0x%016X", expected, actual)
	}
}

func TestCrc64EmptyVector(t *testing.T) {
	AssetEqual(t, 0x0000000000000000, Checksum(nil, 0xD800000000000000, 0x0000000000000000, 0x0000000000000000))
}

// http://www.sunshine2k.de/coding/javascript/crc/crc_js.html
// https://crccalc.com/
func TestNormalized(t *testing.T) {

	input := []byte("abcdefgh")

	t.Log("CRC64_ECMA_182")
	{
		AssetEqual(t, 0x6641AB24513DBCCB, Checksum(input, 0x42F0E1EBA9EA3693, 0x0000000000000000, 0x0000000000000000))
	}

	//t.Log("CRC64_GO_ISO")
	//{
	//	AssetEqual(t, 0x0E21B002D36776C4, Checksum(input, 0x000000000000001B, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, true, true))
	//}

	t.Log("CRC64_WE")
	{
		AssetEqual(t, 0x6512EA66F7F3EAA6, Checksum(input, 0x42F0E1EBA9EA3693, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF))
	}

	//t.Log("CRC64_XZ")
	//{
	//AssetEqual(t, 0x0E21B002D36776C4, Checksum(input, 0x000000000000001B, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF))
	//}
}

func BenchmarkCrcSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		Checksum(smallText, 0x42F0E1EBA9EA3693, 0x0000000000000000, 0x0000000000000000)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		Checksum(largeText, 0x42F0E1EBA9EA3693, 0x0000000000000000, 0x0000000000000000)
	}
}
