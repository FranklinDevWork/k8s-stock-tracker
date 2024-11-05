package models

type LastNDays struct {
	Symbol          string       `json:"symbol"`
	NumberOfDays    int32        `json:"number_of_days"`
	AverageOverDays float64      `json:"average_closing_price"`
	NDaysOfResults  []SeriesData `json:"results"`
}
