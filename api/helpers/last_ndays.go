package helpers

import (
	"errors"
	"strconv"

	"github.com/FranklinDevWork/k8s-stock-tracker/api/models"
)

// extract the last ndays of data from the alphavantage time
// series data
func LastNDaysFromAV(data models.AlphaVantageResponse, nDays int32) models.LastNDays {
	// not really nDays, more N data entries, but for now
	// lets roll with it
	keysForNDays := keySlice(data.TimeSeriesDaily)[:nDays-1]

	timeSeriesData := []models.SeriesData{}
	for _, key := range keysForNDays {
		timeSeriesData = append(timeSeriesData, data.TimeSeriesDaily[key])
	}

	lastNDays := models.LastNDays{}

	lastNDays.NDaysOfResults = timeSeriesData
	lastNDays.NumberOfDays = nDays
	lastNDays.Symbol = data.MetaData.Symbol
	// extract average closing price, if it errors, we will be very cheeky
	if avgClosingPrice, err := averageClosingPrice(timeSeriesData); err == nil {
		lastNDays.AverageOverDays = avgClosingPrice
	}
	return lastNDays
}

func keySlice(data models.TimeSeriesData) []string {
	keys := []string{}
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}

func averageClosingPrice(data []models.SeriesData) (float64, error) {
	sum := 0.0
	for _, series := range data {
		closeValue, err := strconv.ParseFloat(series.Close, 64)
		if err != nil {
			return -1, errors.New("issues parsing closing value as a float")
		}
		sum += closeValue
	}
	return float64(sum) / float64(len(data)), nil
}
