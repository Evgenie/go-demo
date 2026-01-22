package main

import "fmt"

const usdToEuro = 0.94
const usdToRub = 80.0

func getParameters() (float64, string, string) {
	var sourceCurrency string
	var targetCurrency string
	var amount float64

	fmt.Println("__ Конвертер валюты __")
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&sourceCurrency)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	return amount, sourceCurrency, targetCurrency
}

func convertCurrency(amount float64, sourceCurrency string, targetCurrency string) float64 {
	return 0.0
}

func main() {
	const euroToRub = usdToRub / usdToEuro
	fmt.Printf("Course euro to rub equals %.2f\n", euroToRub)
}
