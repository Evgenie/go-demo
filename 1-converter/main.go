package main

import (
	"fmt"
	"slices"
	"strings"
)

type Courses = map[string]map[string]float64

func getSourceCurrency(currencies string) string {
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
	const currencies = "usd/eur/rub"
	currencyArr := strings.Split(currencies, "/")

	sourceCurrency := getSourceCurrency(currencies)
	for !slices.Contains(currencyArr, sourceCurrency) {
		fmt.Println("Ошибка. Повторите ввод")
		sourceCurrency = getSourceCurrency(currencies)
	}
	sourceCurrencyIdx := slices.Index(currencyArr, sourceCurrency)
	suggestCurrencyArr := slices.Delete(currencyArr, sourceCurrencyIdx, sourceCurrencyIdx+1)
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

func convertCurrency(courses *Courses) {
	amount, sourceCurrency, targetCurrency := getParameters()
	converted := amount * (*courses)[sourceCurrency][targetCurrency]

	fmt.Printf("Course %s to %s equals %.2f\n", sourceCurrency, targetCurrency, converted)
}

func main() {
	const usdToEur = 0.94
	const usdToRub = 80.0

	eurToUsd := 1 / usdToEur
	eurToRub := usdToRub / usdToEur
	rubToUsd := 1 / usdToRub
	rubToEur := 1 / eurToRub

	courses := Courses{
		"usd": {
			"eur": usdToEur,
			"rub": usdToRub,
		},
		"eur": {
			"usd": eurToUsd,
			"rub": eurToRub,
		},
		"rub": {
			"usd": rubToUsd,
			"eur": rubToEur,
		},
	}

	fmt.Println("__ Конвертер валюты __")
	convertCurrency(&courses)
}
