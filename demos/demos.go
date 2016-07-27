package demos

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	//"fmt"
)

var (
	demosPath string //The selected path to look for demos
	demos     []Demo //Slice holding all of the demos found in demosPath
)

func init() {
	setPathToDefault()
}

//SetPath sets package variable demosPath, which is used to search for demo files
func SetPath(p string) {
	demosPath = p
}

//setPathToDefault sets the demosPath variable to what should be the default tf directory
//TODO check for Windows
func setPathToDefault() {
	usr, _ := user.Current()
	demosPath = fmt.Sprintf("%v/%v", usr.HomeDir, `.steam/steam/steamapps/common/Team Fortress 2/tf/`)
	fmt.Println(demosPath)
}

//GetDemos calls Walk with demoVisitor, which in turn stores all of the demo
//details found in the package variable demos and then returns that variable.
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
