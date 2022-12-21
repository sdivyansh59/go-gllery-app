package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	a := app.New()
	w := a.NewWindow("Hello World")
	
	root_src := "C:\\Users\\Divyansh Singh\\Pictures"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
    var picsArr [] string

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false {
			extention := strings.Split(file.Name(),".")[1]
			if extention == "png" || extention == "jpeg" {
				picsArr = append(picsArr, root_src+"\\"+file.Name())
			}
		}
	}
//    for _, val := range picsArr {
// 		image := canvas.NewImageFromFile(val)
// 		w.SetContent(image)
//    }

   tabs := container.NewAppTabs(
	container.NewTabItem("Image2", canvas.NewImageFromFile(picsArr[0])),
   )
   for i := 1; i< len(picsArr) ; i++ {
	tabs.Append(container.NewTabItem("Image", canvas.NewImageFromFile(picsArr[i])))
   }
	tabs.SetTabLocation(container.TabLocationLeading)
	w.Resize(fyne.NewSize(600,600))
	w.SetContent(tabs)
	w.ShowAndRun()
}