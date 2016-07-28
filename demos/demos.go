package demos

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	//"fmt"
)

var (
	demosPath string //The selected path to look for demos
	demos     []Demo //Slice holding all of the demos found in demosPath
)

func init() {
	setPathFromFile()
}

//Sorting functions for []Demo

func SortByNameAZ() { sort.Sort(ByName(demos)) }
func SortByNameZA() { sort.Sort(sort.Reverse(ByName(demos))) }

func SortByMapAZ() { sort.Sort(ByMap(demos)) }
func SortByMapZA() { sort.Sort(sort.Reverse(ByMap(demos))) }

func SortByUserAZ() { sort.Sort(ByUser(demos)) }
func SortByUserZA() { sort.Sort(sort.Reverse(ByUser(demos))) }

func SortByDateOldest() { sort.Sort(ByDate(demos)) }
func SortByDateNewest() { sort.Sort(sort.Reverse(ByDate(demos))) }

//ByName implements sort.Interface for []Demo by name
type ByName []Demo

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name() < a[j].Name() }

//ByMap implements sort.Interface for []Demo by map
type ByMap []Demo

func (a ByMap) Len() int           { return len(a) }
func (a ByMap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMap) Less(i, j int) bool { return a[i].MapName() < a[j].MapName() }

//ByUser implements sort.Interface for []Demo by username
type ByUser []Demo

func (a ByUser) Len() int           { return len(a) }
func (a ByUser) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUser) Less(i, j int) bool { return a[i].ClientName() < a[j].ClientName() }

//ByDate implements sort.Interface for []Demo by date
type ByDate []Demo

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date().Before(a[j].Date()) }

//SetPath sets package variable demosPath, which is used to search for demo files
func SetPath(p string) {
	fmt.Println("Setting path: ", p)
	demosPath = p
	demos = demos[:0]
	ioutil.WriteFile("path.data", []byte(p), 0644)
	filepath.Walk(demosPath, demoVisitor)
}

//setPathFromFile reads the path.data file for a saved directory to load
func setPathFromFile() {
	path, err := ioutil.ReadFile("path.data")
	if err == nil {
		SetPath(string(path))
	}
}

//GetDemos calls Walk with demoVisitor, which in turn stores all of the demo
//details found in the package variable demos and then returns that variable.
func GetDemos() []Demo {
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
