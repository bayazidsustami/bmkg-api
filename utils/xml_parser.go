package utils

import (
	"strconv"
	"strings"

	"github.com/bayazidsustami/bmkg-api/models"
	"github.com/beevik/etree"
)

func ParseAllElement(response string) (*models.Weather, error) {
	doc := etree.NewDocument()
	err := doc.ReadFromString(response)
	if err != nil {
		return nil, err
	}

	root := doc.SelectElement("data")
	forecastField := root.SelectElement("forecast")
	issueField := forecastField.SelectElement("issue")

	data := models.Weather{
		TimeStamp: int64(getElementInt(issueField, "timestamp")),
		Areas:     getAreas(forecastField),
	}

	return &data, nil
}

func getAreas(forecastField *etree.Element) []models.Area {
	var items []models.Area
	for _, element := range forecastField.SelectElements("area") {
		coordinate := strings.Replace(getAttrString(element, "coordinate"), " ", ",", -1)

		itemArea := models.Area{
			Id:          rune(getAttrInt(element, "id")),
			Latitude:    getAttrFloat(element, "latitude"),
			Longitude:   getAttrFloat(element, "longitude"),
			Coordinate:  coordinate,
			Type:        getAttrString(element, "type"),
			Region:      getAttrString(element, "region"),
			Level:       rune(getAttrInt(element, "level")),
			Description: getAttrString(element, "description"),
			Domain:      getAttrString(element, "domain"),
			Tags:        getAttrString(element, "tags"),
			Parameter: models.Parameter{
				Humidity:       getHumidity(element),
				MaxHumidity:    getMaxHumidity(element),
				MinHumidity:    getMinHumidity(element),
				Temperature:    getTemperature(element),
				MinTemperature: getMinTemperature(element),
				MaxTemperature: getMaxTemperature(element),
			},
			Name: element.SelectElement("name").Text(),
		}

		items = append(items, itemArea)
	}

	return items
}

func getHumidity(element *etree.Element) models.Humidity {
	elementHummidity := findParameterById(element, "hu")

	var humidityValues []models.HumidityValue
	for _, e := range elementHummidity.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		humidityValue := models.HumidityValue{
			Humidity: rune(getAttrInt(e, "h")),
			DateTime: int64(getAttrInt(e, "datetime")),
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
		}
		humidityValues = append(humidityValues, humidityValue)
	}

	return models.Humidity{
		Id:          getAttrString(elementHummidity, "id"),
		Description: getAttrString(elementHummidity, "description"),
		Type:        getAttrString(elementHummidity, "type"),
		Data:        humidityValues,
	}
}

func getMaxHumidity(element *etree.Element) models.MinMaxHumidity {
	elementMaxHumidity := findParameterById(element, "humax")

	var maxHumidityValues []models.MinMaxHumidityValue
	for _, e := range elementMaxHumidity.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		maxHumidity := models.MinMaxHumidityValue{
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
			DateTime: int64(getAttrInt(e, "datetime")),
		}
		maxHumidityValues = append(maxHumidityValues, maxHumidity)
	}

	return models.MinMaxHumidity{
		Id:          getAttrString(elementMaxHumidity, "id"),
		Description: getAttrString(elementMaxHumidity, "description"),
		Type:        getAttrString(elementMaxHumidity, "type"),
		Data:        maxHumidityValues,
	}
}

func getMinHumidity(element *etree.Element) models.MinMaxHumidity {
	elementMaxHumidity := findParameterById(element, "humin")

	var maxHumidityValues []models.MinMaxHumidityValue
	for _, e := range elementMaxHumidity.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		maxHumidity := models.MinMaxHumidityValue{
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
			DateTime: int64(getAttrInt(e, "datetime")),
		}
		maxHumidityValues = append(maxHumidityValues, maxHumidity)
	}

	return models.MinMaxHumidity{
		Id:          getAttrString(elementMaxHumidity, "id"),
		Description: getAttrString(elementMaxHumidity, "description"),
		Type:        getAttrString(elementMaxHumidity, "type"),
		Data:        maxHumidityValues,
	}
}

func getTemperature(element *etree.Element) models.Temperature {
	elementTemp := findParameterById(element, "t")

	var tempValues []models.TemperatureValue
	for _, e := range elementTemp.SelectElements("timerange") {
		valueElement := e.SelectElement("value")

		tempValue := models.TemperatureValue{
			Humidity: rune(getAttrInt(e, "h")),
			DateTime: int64(getAttrInt(e, "datetime")),
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
		}
		tempValues = append(tempValues, tempValue)
	}

	return models.Temperature{
		Id:          getAttrString(elementTemp, "id"),
		Description: getAttrString(elementTemp, "description"),
		Type:        getAttrString(elementTemp, "type"),
		Data:        tempValues,
	}
}

func getMaxTemperature(element *etree.Element) models.MinMaxTemperature {
	elementMaxTemperature := findParameterById(element, "tmin")

	var maxTemperatureValues []models.MinMaxTemperatureValue
	for _, e := range elementMaxTemperature.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		maxTemperature := models.MinMaxTemperatureValue{
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
			DateTime: int64(getAttrInt(e, "datetime")),
		}
		maxTemperatureValues = append(maxTemperatureValues, maxTemperature)
	}

	return models.MinMaxTemperature{
		Id:          getAttrString(elementMaxTemperature, "id"),
		Description: getAttrString(elementMaxTemperature, "description"),
		Type:        getAttrString(elementMaxTemperature, "type"),
		Data:        maxTemperatureValues,
	}
}

func getMinTemperature(element *etree.Element) models.MinMaxTemperature {
	elementMaxTemperature := findParameterById(element, "tmax")

	var maxTemperatureValues []models.MinMaxTemperatureValue
	for _, e := range elementMaxTemperature.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		maxTemperature := models.MinMaxTemperatureValue{
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
			DateTime: int64(getAttrInt(e, "datetime")),
		}
		maxTemperatureValues = append(maxTemperatureValues, maxTemperature)
	}

	return models.MinMaxTemperature{
		Id:          getAttrString(elementMaxTemperature, "id"),
		Description: getAttrString(elementMaxTemperature, "description"),
		Type:        getAttrString(elementMaxTemperature, "type"),
		Data:        maxTemperatureValues,
	}
}

func findParameterById(element *etree.Element, key string) *etree.Element {
	param := "//parameter[@id='" + key + "']"
	for _, e := range element.FindElements(param) {
		id := e.SelectAttrValue("id", "")
		if id == key {
			return e
		}
	}
	return nil
}

func getElementInt(element *etree.Element, key string) int {
	value, _ := strconv.Atoi(element.SelectElement(key).Text())
	return value
}

func getAttrInt(element *etree.Element, key string) int {
	attrValue := element.SelectAttrValue(key, "0")
	value, _ := strconv.Atoi(attrValue)
	return value
}

func getAttrFloat(element *etree.Element, key string) float64 {
	attrValue := element.SelectAttrValue(key, "0")
	value, _ := strconv.ParseFloat(attrValue, 64)
	return value
}

func getAttrString(element *etree.Element, key string) string {
	return element.SelectAttrValue(key, "")
}
