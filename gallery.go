package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func GalleryApp(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.Size{800, 600});

	root_src := "C:\\Users\\Ganesh\\Pictures"
	files, err := ioutil.ReadDir(root_src);
	if err != nil {
		log.Fatal(err)
	}
	var picsArray []string;
	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1];
			if extension == "png" || extension == "jpeg" {
				picsArray = append(picsArray, root_src + "\\" + file.Name());
			}
		}
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Image 1", canvas.NewImageFromFile(picsArray[0])),
	)
	for i := 1; i < len(picsArray); i++ {
		tabs.Append(container.NewTabItem("Image "+ strconv.Itoa(i + 1), canvas.NewImageFromFile(picsArray[i])))
	}
	tabs.SetTabLocation(container.TabLocationTrailing);
	
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs),)
	// w.SetContent(image);
	w.Show()
}
