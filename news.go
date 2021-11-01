package main

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var num_article int = 1
var news News

func showNews(w fyne.Window) {
	URL := "https://gnews.io/api/v4/top-headlines?token=6c2ea826d3ff4ad963b94dfb3ba41220&lang=en&max=100"
	//API
	res, err := http.Get(URL)
	if(err != nil){
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	news, _ = UnmarshalNews(body)

	// num_article = int(news.TotalArticles)
	// label2 := widget.NewLabel(fmt.Sprintf("No of articles:%d",
	// 	news.TotalArticles))

	//show title
	label3 := widget.NewLabel(news.Articles[0].Title)
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label3.Wrapping = fyne.TextWrapBreak
	
	// show articles
	entry1 := widget.NewLabel(news.Articles[1].Description)
	// entry1.MultiLine = true
	entry1.Wrapping = fyne.TextWrapBreak
	entry1.Resize(fyne.NewSize(100, 100))
	// label4 := widget.NewLabel(fmt.Sprintf("Article: %s", news.Articles[0].Content))
	// label4.Resize(fyne.NewSize(200, 200))
	btn := widget.NewButton("Next", func() {
		num_article += 1
		label3.Text = news.Articles[num_article].Title
		// label3.TextStyle = fyne.TextStyle{Bold: true}
		// label3.Wrapping = fyne.TextWrapBreak
		entry1.Text = news.Articles[num_article].Description
		// entry1.Resize(fyne.NewSize(100, 300))
		// label4.Text = news.Articles[num_article].Content
		// label4.Resize(fyne.NewSize(40, 60))
		label3.Refresh()
		entry1.Refresh()
		// label4.Refresh()
	})
	label1 := canvas.NewText("News App", color.White)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextStyle = fyne.TextStyle{Bold: true}
	img := canvas.NewImageFromFile("news.png")
	img.Resize(fyne.NewSize(200, 200))
	img.FillMode = canvas.ImageFillOriginal
	newsContainer := container.NewVBox(label1, label3, entry1, btn)
	newsContainer.Resize(fyne.NewSize(300, 300))
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, newsContainer))
	w.Show()
}


// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()
func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	TotalArticles int64     `json:"totalArticles"`
	Articles      []Article `json:"articles"`
}
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
	Source      Source `json:"source"`
}
type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
