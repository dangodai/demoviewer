package demos

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	//"fmt"
)

var (
	demosPath string
	demos     []Demo
)

func init() {
	SetPathToDefault()
}

//Maybe error check and make sure the path is valid?
func SetPath(p string) {
	demosPath = p
}

func SetPathToDefault() {
	usr, _ := user.Current()
	demosPath = fmt.Sprintf("%v/%v", usr.HomeDir, `.steam/steam/steamapps/common/Team Fortress 2/tf/`)
	fmt.Println(demosPath)
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
