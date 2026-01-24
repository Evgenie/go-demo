package main

import (
	"fmt"
	"slices"
	"strings"
)

const currencies = "usd/eur/rub"
const usdToEur = 0.94
const usdToRub = 80.0
var currencyArr = strings.Split(currencies, "/")
var eurToUsd = 1 / usdToEur
var eurToRub = usdToRub / usdToEur
var rubToUsd = 1 / usdToRub
var rubToEur = 1 / eurToRub

func getSourceCurrency() string {
	var sourceCurrency string
	fmt.Printf("Введите исходную валюту(%s): ", currencies)
	fmt.Scan(&sourceCurrency)
	return sourceCurrency
}
func getTargetCurrency(suggestCurrencies string) string {
	var targetCurrency string
	fmt.Printf("Введите целевую валюту(%s): ", suggestCurrencies)
	fmt.Scan(&targetCurrency)
	return targetCurrency
}
func getAmount() (float64, error) {
	var amount float64
	fmt.Printf("Введите сумму: ")
	_, error := fmt.Scan(&amount)
	return amount, error
}

func getParameters() (float64, string, string) {
	sourceCurrency := getSourceCurrency()
	for !slices.Contains(currencyArr, sourceCurrency) {
		fmt.Println("Ошибка. Повторите ввод")
		sourceCurrency = getSourceCurrency()
	}
	sourceCurrencyIdx := slices.Index(currencyArr, sourceCurrency)
	suggestCurrencyArr := slices.Delete(currencyArr, sourceCurrencyIdx, sourceCurrencyIdx + 1)
	suggestCurrencies := strings.Join(suggestCurrencyArr, "/")
	targetCurrency := getTargetCurrency(suggestCurrencies)
	for !slices.Contains(suggestCurrencyArr, targetCurrency) {
		fmt.Println("Ошибка. Повторите ввод")
		targetCurrency = getTargetCurrency(suggestCurrencies)
	}
	amount, err := getAmount()
	for amount <= 0 || err != nil {
		fmt.Println("Ошибка. Повторите ввод")
		amount, err = getAmount()
	}
	
	return amount, sourceCurrency, targetCurrency
}

func convertCurrency(amount float64, sourceCurrency string, targetCurrency string) {
	var converted float64
	switch sourceCurrency {
		case "usd":
			if targetCurrency == "rub" {
				converted = amount * usdToRub
				} else {
				converted = amount * usdToEur
			}
		case "eur":
			if targetCurrency == "rub" {
				converted = amount * eurToRub
				} else {
				converted = amount * eurToUsd
			}
		default:
			if targetCurrency == "usd" {
				converted = amount * rubToUsd
			} else {
				converted = amount * rubToEur
			}
	}
	fmt.Printf("Course %s to %s equals %.2f\n",sourceCurrency, targetCurrency,  converted)
}

func main() {
	fmt.Println("__ Конвертер валюты __")
	convertCurrency(getParameters())
}
