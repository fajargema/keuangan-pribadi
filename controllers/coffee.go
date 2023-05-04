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
    RealtimeCurrencyExchangeRate struct {
        FromCurrencyCode string `json:"1. From_Currency Code"`
        FromCurrencyName string `json:"2. From_Currency Name"`
        ToCurrencyCode   string `json:"3. To_Currency Code"`
        ToCurrencyName   string `json:"4. To_Currency Name"`
        ExchangeRate     string `json:"5. Exchange Rate"`
        LastRefreshed    string `json:"6. Last Refreshed"`
        TimeZone         string `json:"7. Time Zone"`
        BidPrice         string `json:"8. Bid Price"`
        AskPrice         string `json:"9. Ask Price"`
    } `json:"Realtime Currency Exchange Rate"`
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

	rpResp, _ := http.Get("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=GBP&to_currency=IDR&apikey=R5KS7LRJO82IIQ55")
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
		rp, _ := strconv.ParseFloat(rpData.RealtimeCurrencyExchangeRate.ExchangeRate, 64)
		res := valueFloat * rp
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