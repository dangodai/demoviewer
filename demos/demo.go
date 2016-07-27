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

//Demo represents all the information gathered from a demo file, demo header,
//and demo json file
type Demo struct {
	info               os.FileInfo
	header             DemoHeader
	events             []Event
	jsonPath, demoPath string
}

//Name returns demo's filename
func (d *Demo) Name() string {
	return d.info.Name()
}

//MapName returns the map the demo was played on
func (d *Demo) MapName() string {
	return d.header.GetMapName()
}

//ServerName returns the server IP:Port, or the STV name from the demo
func (d *Demo) ServerName() string {
	return d.header.GetServerName()
}

//ClientName returns the name of the client that recorded the demo (ie their
//steam display name at the time the demo was recorded)
func (d *Demo) ClientName() string {
	return d.header.GetClientName()
}

//Type returns whether the demo is a POV or STV demo, based on the ServerName
//The check is rather simple and could probably use from regexp
func (d *Demo) Type() string {
	if d.ServerName() == "localhost" || len(strings.Split(d.ServerName(), ":")) == 2 {
		return "POV"
	}
	return "STV"
}

//Path returns the absolute path to the demo file
func (d *Demo) Path() string {
	return d.demoPath
}

//Date returns the last modified time of the demo file, reported by the OS
func (d *Demo) Date() time.Time {
	return d.info.ModTime()
}

//PathInTFFolder returns path to the demo from within the TF folder
//(Since TF2 only searches directly in the TF folder unless told otherwise)
//A little bit hackish
func (d *Demo) PathInTFFolder() string {
	//Make sure we don't crash if they don't choose the tf folder properly
	temp := strings.Split(d.Path(), "/tf/")
	return temp[len(temp)-1]
}

//Events returns a slice of Event objects parsed from any json file corresponding
//to the demo
func (d *Demo) Events() []Event {
	return d.events
}

//Play attempts to launch TF2 with the +playdemo launch option, specifying the
//demo to be played. Very hackish on windows but the best I could come up with
func (d *Demo) Play() {
	//Steam isn't in PATH on windows systems, have to specify steam path
	cmd := exec.Command("steam", "-applaunch", "440", "+playdemo", d.PathInTFFolder())
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", "", `C:\Program Files (x86)\Steam\Steam.exe`,
			"-applaunch", "440", "+playdemo", d.PathInTFFolder())
	}
	cmd.Start()
}

//Delete delete the demo file and any matching json file from the computer
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

//EventResponse is a container for Event objects, used to easily parse the json
//into objects
type EventResponse struct {
	Events []Event `json:"events"`
}

//Event represents events found in a demo's json file
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
