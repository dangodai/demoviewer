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
	events             []Event
	jsonPath, demoPath string
}

func (d *Demo) Name() string {
	return d.info.Name()
}

func (d *Demo) Path() string {
	return d.demoPath
}

func (d *Demo) Date() time.Time {
	return d.info.ModTime()
}

//Returns the path to the demo from within the TF folder
//(Since TF2 only searches directly in the TF folder unless told otherwise)
//A little bit hackish
func (d *Demo) PathInTFFolder() string {
	return strings.Replace(d.Path(), demosPath+string(os.PathSeparator), "", 1)
}

func (d *Demo) Events() []Event {
	return d.events
}

func (d *Demo) Play() {
	//Steam isn't in PATH on windows systems, have to specify steam path
	command := "steam"
	if runtime.GOOS == "windows" {
		command = `C:\Program Files(x86)\Steam\Steam.exe`
	}
	cmd := exec.Command(command, "-applaunch", "440", "+playdemo", d.PathInTFFolder())
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
