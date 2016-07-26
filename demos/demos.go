package demos

import (
	"os"
	"path/filepath"
	//"fmt"
)

var (
	demosPath string
	demos     []Demo
)

//Maybe error check and make sure the path is valid?
func SetPath(p string) {
	demosPath = p
}

func GetDemos() []Demo {
	demos = demos[:0]
	filepath.Walk(demosPath, demoVisitor)
	return demos
}

func demoVisitor(path string, info os.FileInfo, err error) error {
	if filepath.Ext(path) != ".dem" {
		return nil
	}

	//Look for the matching json file
	jsonPath := path[:len(path)-4] + ".json"
	events := []Event{}
	_, e := os.Open(jsonPath)
	if e == nil {
		events = getEvents(jsonPath)
	} else {
		events = nil
		jsonPath = ""
	}

	demos = append(demos, Demo{
		info:     info,
		header:   ReadHeaderFromFile(path),
		events:   events,
		jsonPath: jsonPath,
		demoPath: path,
	})

	return nil
}
