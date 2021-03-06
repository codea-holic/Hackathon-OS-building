package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func weatherApp(w fyne.Window) {

	// API PART
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=pune&APPID=c8c16d5d3b3508247e1544f634d1370c")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}

	img := canvas.NewImageFromFile("weather.png")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("Weather App", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}

	label2 := canvas.NewText(fmt.Sprintf("country %s", weather.Sys.Country), color.White)
	label3 := canvas.NewText(fmt.Sprintf("Wind Speed %.2f", weather.Wind.Speed), color.White)
	label4 := canvas.NewText(fmt.Sprintf("Temperature %2f", weather.Main.Temp), color.White)
	label5 := canvas.NewText(fmt.Sprintln("Humidity ", weather.Main.Humidity), color.White)
	label6 := canvas.NewText(fmt.Sprintln("Pressure ", weather.Main.Pressure), color.White)

	weatherContainer := container.NewVBox(
		label1,
		img,
		label2,
		label3,
		label4,
		label5,
		label6,
		container.NewGridWithColumns(1),
	)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, weatherContainer))
	w.Show()
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
