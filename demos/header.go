package demos

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	HeaderLength = 1072
)

type DemoHeader struct {
	Header          [8]byte
	DemoProtocol    uint32
	NetworkProtocol uint32
	ServerName      [260]byte
	ClientName      [260]byte
	MapName         [260]byte
	GameDirectory   [260]byte
	PlaybackTime    float32
	Ticks           uint32
	Frames          uint32
	SignOnLength    uint32
}

func ReadHeaderFromFile(path string) DemoHeader {
	header := DemoHeader{}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer f.Close()

	/*rawBytes := make([]byte, HeaderLength)
	f.Read(rawBytes)

	buffer := bytes.NewBuffer(rawBytes)*/

	binary.Read(f, binary.LittleEndian, &header)
	return header
}

func (h *DemoHeader) GetServerName() string {
	return string(h.ServerName[:])
}

func (h *DemoHeader) GetClientName() string {
	return string(h.ClientName[:])
}

func (h *DemoHeader) GetMapName() string {
	return string(h.MapName[:])
}

func (h *DemoHeader) GetPlaybackTime() float32 {
	return h.PlaybackTime
}
