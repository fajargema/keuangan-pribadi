package controllers

import (
	"encoding/json"
	"keuangan-pribadi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leekchan/accounting"
)

func GetCoffeePrice(c echo.Context) error {
    // buat request ke API
    resp, err := http.Get("https://www.alphavantage.co/query?function=COFFEE&interval=monthly&apikey=R5KS7LRJO82IIQ55")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
    }
    defer resp.Body.Close()

    // parse response dari API
    var data struct {
        Name     string          `json:"name"`
        Interval string          `json:"interval"`
        Unit     string          `json:"unit"`
        Data     []json.RawMessage `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
    }

    // konversi setiap elemen data menjadi objek CoffeePrice
    var coffeePrices []models.CoffeePrice
    for _, elem := range data.Data {
        var coffeePrice models.CoffeePrice
        if err := json.Unmarshal(elem, &coffeePrice); err != nil {
            return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
        }
		valueFloat, _ := strconv.ParseFloat(coffeePrice.Value, 64)
		res := valueFloat * 18000
		// coffeePrice.Value = strconv.FormatFloat(res, 'f', 2, 64)
		ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ".", Decimal: ","}
		coffeePrice.Value = ac.FormatMoney(res)
        coffeePrices = append(coffeePrices, coffeePrice)
    }

    // buat response dengan format JSON
    response := models.CoffeePriceResponse{
        Name:     data.Name,
        Interval: data.Interval,
        Unit:     data.Unit,
        Data:     coffeePrices,
    }
    return c.JSON(http.StatusOK, response)
}