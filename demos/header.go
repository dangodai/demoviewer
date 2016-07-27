package demos

import (
	"encoding/binary"
	"log"
	"os"
)

const (
	headerLength = 1072 //Total length of the header in bytes
)

//DemoHeader is a structure matching a .dem file's header structure
//See https://developer.valvesoftware.com/wiki/DEM_Format
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

//ReadHeaderFromFile opens the file specified by path and copies the header bytes
//into a DemoHeader object. The created DemoHeader is then returned.
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

//GetServerName is a convenience function to get the server name from a header
//as a string instead of a byte slice
func (h *DemoHeader) GetServerName() string {
	return string(h.ServerName[:])
}

//GetClientName is a convenience function to get the client name from a header
//as a string instead of a byte slice
func (h *DemoHeader) GetClientName() string {
	return string(h.ClientName[:])
}

//GetMapName is a convenience function to get the map name from a header
//as a string instead of a byte slice
func (h *DemoHeader) GetMapName() string {
	return string(h.MapName[:])
}

//GetPlaybackTime is a convenience function to get the PlaybackTime from a header
//Actually I don't really need this.
func (h *DemoHeader) GetPlaybackTime() float32 {
	return h.PlaybackTime
}
