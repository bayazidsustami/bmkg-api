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
				Humidity:    getHumidity(element),
				MaxHumidity: getMaxHumidity(element),
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

func getMaxHumidity(element *etree.Element) models.MaxHumidity {
	elementMaxHumidity := findParameterById(element, "humax")

	var maxHumidityValues []models.MaxHumidityValue
	for _, e := range elementMaxHumidity.SelectElements("timerange") {
		valueElement := e.SelectElement("value")
		maxHumidity := models.MaxHumidityValue{
			Value:    rune(getElementInt(e, "value")),
			Unit:     getAttrString(valueElement, "unit"),
			DateTime: int64(getAttrInt(e, "datetime")),
		}
		maxHumidityValues = append(maxHumidityValues, maxHumidity)
	}

	return models.MaxHumidity{
		Id:          getAttrString(elementMaxHumidity, "id"),
		Description: getAttrString(elementMaxHumidity, "description"),
		Type:        getAttrString(elementMaxHumidity, "type"),
		Data:        maxHumidityValues,
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
