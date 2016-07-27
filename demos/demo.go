package demos

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Demo struct {
	info               os.FileInfo
	header             DemoHeader
	events             []Event
	jsonPath, demoPath string
}

func (d *Demo) Name() string {
	return d.info.Name()
}

func (d *Demo) MapName() string {
	return d.header.GetMapName()
}

func (d *Demo) ServerName() string {
	return d.header.GetServerName()
}

func (d *Demo) ClientName() string {
	return d.header.GetClientName()
}

func (d *Demo) Path() string {
	return d.demoPath
}

func (d *Demo) Date() time.Time {
	return d.info.ModTime()
}

//PathInTFFolder returns path to the demo from within the TF folder
//(Since TF2 only searches directly in the TF folder unless told otherwise)
//A little bit hackish
func (d *Demo) PathInTFFolder() string {
	return strings.Split(d.Path(), "/tf/")[1]
}

func (d *Demo) Events() []Event {
	return d.events
}

func (d *Demo) Play() {
	//Steam isn't in PATH on windows systems, have to specify steam path
	cmd := exec.Command("steam", "-applaunch", "440", "+playdemo", d.PathInTFFolder())
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "", `C:\Program Files (x86)\Steam\Steam.exe`,
			"-applaunch", "440", "+playdemo", d.PathInTFFolder())
	}
	cmd.Start()
}

func (d *Demo) Delete() bool {
	if d.jsonPath != "" {
		if err := os.Remove(d.jsonPath); err != nil {
			return false
		}
	}
	if err := os.Remove(d.demoPath); err != nil {
		return false
	}

	return true
}

type EventResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Tick  int    `json:"tick"`
}

func getEvents(path string) []Event {
	var events EventResponse
	data, _ := ioutil.ReadFile(path)
	json.Unmarshal(data, &events)

	if len(events.Events) != 0 {
		return events.Events
	}
	return nil

}
