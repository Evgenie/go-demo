package main

import "fmt"

func main() {
	const usdToEuro = 0.94
	const usdToRub = 80.0
	const euroToRub = usdToRub / usdToEuro
	fmt.Printf("Course euro to rub equals %.2f\n", euroToRub)
}
