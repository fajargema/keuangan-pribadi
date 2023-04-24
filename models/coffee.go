package models

type CoffeePrice struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}

type CoffeePriceResponse struct {
	Name     string        `json:"name"`
	Interval string        `json:"interval"`
	Unit     string        `json:"unit"`
	Data     []CoffeePrice `json:"data"`
}