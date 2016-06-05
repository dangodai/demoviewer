package demos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
  "os/exec"
  "strings"
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

//Returns the path to the demo from within the TF folder
//(Since TF2 only searches directly in the TF folder unless told otherwise)
//A little bit hackish
func (d *Demo) PathInTFFolder() string {
  return strings.Split(d.Path(), "Team Fortress 2/tf/")[1]
}

func (d *Demo) Events() []Event {
	return d.events
}

func (d *Demo) Play() {
  cmd := exec.Command("steam", "-applaunch", "440", "+playdemo", d.PathInTFFolder())
  err := cmd.Start()
  fmt.Println(err)
}

type EventResponse struct {
  Events []Event `json:"events"`
}

type Event struct {
	Name string `json:"name"`
  Value string `json:"value"`
  Tick int `json:"tick"`
}

func getEvents(path string) []Event {
	var events EventResponse
	data, _ := ioutil.ReadFile(path)
	json.Unmarshal(data, &events)

	return events.Events
}
