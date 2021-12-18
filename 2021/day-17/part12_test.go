package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPacketReader1(t *testing.T) {
	hex := []byte("D2FE28")
	packet := NewPacketReader(hex).Read()
	require.Equal(t, 2021, packet.LiteralValue)
}

func TestPacketReader2(t *testing.T) {
	hex := []byte("38006F45291200")
	packet := NewPacketReader(hex).Read()
	require.Len(t, packet.SubPackets, 2)
	require.Equal(t, 10, packet.SubPackets[0].LiteralValue)
	require.Equal(t, 20, packet.SubPackets[1].LiteralValue)
}

func TestPacketReader3(t *testing.T) {
	hex := []byte("EE00D40C823060")
	packet := NewPacketReader(hex).Read()
	require.Len(t, packet.SubPackets, 3)
	require.Equal(t, 1, packet.SubPackets[0].LiteralValue)
	require.Equal(t, 2, packet.SubPackets[1].LiteralValue)
	require.Equal(t, 3, packet.SubPackets[2].LiteralValue)
}

func TestPacketReader4(t *testing.T) {
	hex := []byte("8A004A801A8002F478")
	packet := NewPacketReader(hex).Read()
	require.Equal(t, 4, packet.Version)
	require.Equal(t, 1, packet.SubPackets[0].Version)
	require.Equal(t, 5, packet.SubPackets[0].SubPackets[0].Version)
	require.Equal(t, 6, packet.SubPackets[0].SubPackets[0].SubPackets[0].Version)
	require.Equal(t, 16, VersionSum(packet))
}

func TestPacketReader5(t *testing.T) {
	for _, tc := range []struct {
		hex    string
		expect int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	} {
		packet := NewPacketReader([]byte(tc.hex)).Read()
		require.Equal(t, tc.expect, Value(packet), tc.hex)
	}
}
