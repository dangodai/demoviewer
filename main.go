package main

import (
	//"os"
	"fmt"

	"github.com/dangodai/demoviewer/demos"
)

func main() {
	demos.SetPath("/home/yung/.local/share/Steam/steamapps/common/Team Fortress 2/tf/")
	dems := demos.GetDemos()
  for _, d := range dems {
    fmt.Println("-----------------------")
    fmt.Println("Demo Name:", d.Name())
    fmt.Println("Demo Path:", d.Path())
    fmt.Println("Events:", d.Events())
  }

  //dems[6].Play()
}
