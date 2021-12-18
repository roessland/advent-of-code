package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"

	. "github.com/roessland/gopkg/mathutil"
)

type Packet struct {
	Version      int
	TypeID       int
	LiteralValue int64
	SubPackets   []Packet
}

func main() {
	packet := ReadInput()
	fmt.Println("Part 1:", VersionSum(packet))
	fmt.Println("Part 2:", Value(packet))
}

func VersionSum(packet Packet) int {
	versionSum := packet.Version
	for _, subPacket := range packet.SubPackets {
		versionSum += VersionSum(subPacket)
	}
	return versionSum
}

func Value(packet Packet) int64 {
	switch packet.TypeID {
	case 0:
		var sum int64 = 0
		for _, subPacket := range packet.SubPackets {
			sum += Value(subPacket)
		}
		return sum
	case 1:
		var prod int64 = 1
		for _, subPacket := range packet.SubPackets {
			prod *= Value(subPacket)
		}
		return prod
	case 2:
		var min int64 = math.MaxInt64
		for _, subPacket := range packet.SubPackets {
			min = MinInt64(min, Value(subPacket))
		}
		return min
	case 3:
		var max int64 = math.MinInt64
		for _, subPacket := range packet.SubPackets {
			max = MaxInt64(max, Value(subPacket))
		}
		return max
	case 4:
		return packet.LiteralValue
	case 5:
		if Value(packet.SubPackets[0]) > Value(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 6:
		if Value(packet.SubPackets[0]) < Value(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	case 7:
		if Value(packet.SubPackets[0]) == Value(packet.SubPackets[1]) {
			return 1
		} else {
			return 0
		}
	}
	panic(fmt.Sprintf("unknown type id %d", packet.TypeID))
}

func ReadInput() Packet {
	asciiHex, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	packet := NewPacketReader(asciiHex).Read()

	return packet
}

type Bit uint8

type Nibble uint8

type Bits []Bit

func (bits Bits) AsInt() int {
	var n int
	for i := 0; i < len(bits); i++ {
		n = n<<1 + int(bits[i])
	}
	return n
}

func (bits Bits) AsInt64() int64 {
	var n int64
	for i := 0; i < len(bits); i++ {
		n = n<<1 + int64(bits[i])
	}
	return n
}

var hexMapping = map[byte]Nibble{
	'0': 0b0000,
	'1': 0b0001,
	'2': 0b0010,
	'3': 0b0011,
	'4': 0b0100,
	'5': 0b0101,
	'6': 0b0110,
	'7': 0b0111,
	'8': 0b1000,
	'9': 0b1001,
	'A': 0b1010,
	'B': 0b1011,
	'C': 0b1100,
	'D': 0b1101,
	'E': 0b1110,
	'F': 0b1111,
}

type PacketReader struct {
	bits Bits
}

func NewPacketReader(asciiHex []byte) *PacketReader {
	var bits Bits
	for _, c := range asciiHex {
		b := hexMapping[c]
		for shift := 3; shift >= 0; shift-- {
			if b&(1<<shift) != 0 {
				bits = append(bits, 1)
			} else {
				bits = append(bits, 0)
			}
		}
	}
	return &PacketReader{bits: bits}
}

func (r *PacketReader) Read() Packet {
	packet, _ := r.readPacket()
	return packet
}

func (r *PacketReader) readPacket() (Packet, int) {
	var initialLen = len(r.bits)

	var packet Packet
	packet.Version, _ = r.readVersion()
	packet.TypeID, _ = r.readTypeID()

	if packet.TypeID == 4 {
		packet.LiteralValue, _ = r.readLiteralValue()
	} else if packet.TypeID != 4 {
		packet.SubPackets, _ = r.readOperator()
	}

	return packet, initialLen - len(r.bits)
}

func (r *PacketReader) readVersion() (int, int) {
	var initialLen = len(r.bits)
	version := r.bits[0:3].AsInt()
	r.bits = r.bits[3:]
	return version, initialLen - len(r.bits)
}

func (r *PacketReader) readTypeID() (int, int) {
	var initialLen = len(r.bits)
	typeID := r.bits[0:3].AsInt()
	r.bits = r.bits[3:]
	return typeID, initialLen - len(r.bits)
}

func (r *PacketReader) readLiteralValue() (int64, int) {
	var initialLen = len(r.bits)

	var literalBits Bits
	for {
		done := false
		if r.bits[0] == 0 {
			done = true
		}
		literalBits = append(literalBits, r.bits[1:5]...)
		r.bits = r.bits[5:]
		if done {
			break
		}
	}
	return literalBits.AsInt64(), initialLen - len(r.bits)
}

func (r *PacketReader) readOperator() ([]Packet, int) {
	var initialLen = len(r.bits)

	lengthTypeId := r.bits[0:1].AsInt()
	r.bits = r.bits[1:]

	var subPackets []Packet
	if lengthTypeId == 0 {
		remainingBitsToRead := r.bits[0:15].AsInt()
		r.bits = r.bits[15:]
		for remainingBitsToRead > 0 {
			subMsg, bitsRead := r.readPacket()
			subPackets = append(subPackets, subMsg)
			remainingBitsToRead -= bitsRead
		}
	} else if lengthTypeId == 1 {
		remainingSubPacketsToRead := r.bits[0:11].AsInt()
		r.bits = r.bits[11:]
		for remainingSubPacketsToRead > 0 {
			subMsg, _ := r.readPacket()
			subPackets = append(subPackets, subMsg)
			remainingSubPacketsToRead--
		}
	}

	return subPackets, initialLen - len(r.bits)
}
