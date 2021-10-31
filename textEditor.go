package main

import (
	// "io/ioutil"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	// "fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func textEditor() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 500))
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel(" My Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add new file", func() {
		content.Add(widget.NewLabel("New File " + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text ... ")
	input.Resize(fyne.NewSize(800, 800))

	saveBtn := widget.NewButton("Save text file", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName("New File " + strconv.Itoa(count) + ".txt")
		saveFileDialog.Show()
	})


	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", readData)
				veiwData := widget.NewMultiLineEntry()
				veiwData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))

				w.SetContent(container.NewScroll(veiwData))
				w.Resize(fyne.NewSize(500, 500))
				w.Show();
			}, w)

			openFileDialog.SetFilter(storage.NewExtensionFileFilter([] string{".txt"}));
			openFileDialog.Show()
	})

	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w = myApp.NewWindow("Text Editor");
	w.Resize(fyne.NewSize(500, 280));

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, content),
	)
	w.Show()
}
