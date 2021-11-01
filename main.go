package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("My Virtual OS")
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var DeskBtn fyne.Widget
var img fyne.CanvasObject
var panelContent *fyne.Container

func main() {
	myApp.Settings().Theme()
	img = canvas.NewImageFromFile("C:\\Users\\Ganesh\\Desktop\\VirtualOS\\Desktop.jpg")

	sunIcon, err := os.Open(".\\Icons\\sun.png")
	if err != nil {
		log.Fatal(err)
	}

	sunReader := bufio.NewReader(sunIcon)

	sun, err := ioutil.ReadAll(sunReader)
	if err != nil {
		log.Fatal(err)
	}

	btn1 = widget.NewButtonWithIcon("Weather App", fyne.NewStaticResource("icon", sun), func() {
		weatherApp(myWindow)
	})

	calcIcon, err := os.Open(".\\Icons\\calculator.png")
	if err != nil {
		log.Fatal(err)
	}

	calcReader := bufio.NewReader(calcIcon)

	calc, err := ioutil.ReadAll(calcReader)
	if err != nil {
		log.Fatal(err)
	}

	btn2 = widget.NewButtonWithIcon("Calculator", fyne.NewStaticResource("icon", calc), func() {
		showCalc()
	})

	photoIcon , err := os.Open(".\\Icons\\photos.png")
	if err != nil {
		log.Fatal(err)
	}

	photoReader := bufio.NewReader(photoIcon)

	photo, err := ioutil.ReadAll(photoReader)
	if err != nil {
		log.Fatal(err)
	}

	btn3 = widget.NewButtonWithIcon("Gallery", fyne.NewStaticResource("icon", photo), func() {
		GalleryApp(myWindow)
	})

	textIcon , err := os.Open(".\\Icons\\text.png")
	if err != nil {
		log.Fatal(err)
	}

	textReader := bufio.NewReader(textIcon)

	text, err := ioutil.ReadAll(textReader)
	if err != nil {
		log.Fatal(err)
	}

	btn4 = widget.NewButtonWithIcon("Text Editor", fyne.NewStaticResource("icon", text), func() {
		textEditor()
	})

	homeIcon, err := os.Open(".\\Icons\\home.png")
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(homeIcon)

	home, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	DeskBtn = widget.NewButtonWithIcon("Home", fyne.NewStaticResource("icon", home), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(1,
		container.NewGridWithColumns(3, DeskBtn, btn1, btn2),
		container.NewGridWithColumns(3, btn3, btn4),
	)))

	myWindow.Resize(fyne.NewSize(1280, 720))

	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()
}

/*
iconFile, err := os.Open("C:/path/to/res/icon.png")
if err != nil {
	log.Fatal(err)
}

r := bufio.NewReader(iconFile)

b, err := ioutil.ReadAll(r)
if err != nil {
	log.Fatal(err)
}

btn := widget.NewButtonWithIcon("Browse", fyne.NewStaticResource("icon", b), func() {
	// do something
})
*/
