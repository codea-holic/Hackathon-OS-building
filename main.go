package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("My OS");

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject;

var panelContent *fyne.Container

func main(){
	myApp.Settings().Theme(theme.DarkTheme())
	img = canvas.NewImageFromFile("C:\\Users\\Ganesh\\Desktop\\VirtualOS\\Desktop.jpg")
	
	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		weatherApp(myWindow)
	})
}
