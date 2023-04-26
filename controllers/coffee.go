package controllers

import (
	"encoding/json"
	"keuangan-pribadi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/leekchan/accounting"
)

type CurrencyRatesResponse struct {
    Success   bool              `json:"success"`
    Timestamp int64             `json:"timestamp"`
    Base      string            `json:"base"`
    Date      string            `json:"date"`
    Rates     map[string]float64 `json:"rates"`
}

func GetCurrencyRates(c echo.Context) error {
    resp, err := http.Get("https://api.apilayer.com/exchangerates_data/latest?symbols=idr&base=gbp")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, nil)
    }
    defer resp.Body.Close()

    var data CurrencyRatesResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return c.JSON(http.StatusInternalServerError, nil)
    }

    return c.JSON(http.StatusOK, data.Rates["IDR"])
}

func GetCoffeePrice(c echo.Context) error {
    resp, err := http.Get("https://www.alphavantage.co/query?function=COFFEE&interval=monthly&apikey=R5KS7LRJO82IIQ55")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
    }
    defer resp.Body.Close()

    var data struct {
        Name     string          `json:"name"`
        Interval string          `json:"interval"`
        Unit     string          `json:"unit"`
        Data     []json.RawMessage `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
    }

	rpResp, _ := http.Get("https://api.apilayer.com/exchangerates_data/latest?symbols=idr&base=gbp")
	var rpData CurrencyRatesResponse
    if err := json.NewDecoder(rpResp.Body).Decode(&rpData); err != nil {
        return c.JSON(http.StatusInternalServerError, nil)
    }

    var coffeePrices []models.CoffeePrice
    for _, elem := range data.Data {
        var coffeePrice models.CoffeePrice
        if err := json.Unmarshal(elem, &coffeePrice); err != nil {
            return c.JSON(http.StatusInternalServerError, models.CoffeePriceResponse{})
        }

		valueFloat, _ := strconv.ParseFloat(coffeePrice.Value, 64)
		res := valueFloat * rpData.Rates["IDR"]
		ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ".", Decimal: ","}
		
		coffeePrice.Value = ac.FormatMoney(res)
        coffeePrices = append(coffeePrices, coffeePrice)
    }

    response := models.CoffeePriceResponse{
        Name:     data.Name,
        Interval: data.Interval,
        Unit:     data.Unit,
        Data:     coffeePrices,
    }
    return c.JSON(http.StatusOK, response)
}