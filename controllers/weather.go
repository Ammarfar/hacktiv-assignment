package controllers

import (
	"assignment/models"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const fileJSON = "public/weather.json"

func IndexWeather(c *gin.Context) {

	// random request setter
	weather := models.Weather{
		Wind:  uint(rand.Intn(21) + 1),
		Water: uint(rand.Intn(14) + 1),
	}

	// json encoding
	file, _ := json.MarshalIndent(weather, "", " ")

	// setter
	_ = ioutil.WriteFile(fileJSON, file, 0644)

	// getter
	file, _ = ioutil.ReadFile(fileJSON)

	// json decoding
	data := models.Weather{}
	_ = json.Unmarshal(file, &data)

	// view rendering
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"wind":        data.Wind,
		"water":       data.Water,
		"windStatus":  data.GetStatusWind(),
		"waterStatus": data.GetStatusWater(),
	})
}
