package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var operations = []string{
	"AVG",
	"SUM",
	"MED",
}

var operationsStr = strings.Join(operations, "/")

func getOperation() string {
	var operation string
	fmt.Printf("Введите операцию %s: ", operationsStr)
	fmt.Scan(&operation)
	operation = strings.ToUpper(operation)
	for !slices.Contains(operations, operation) {
		fmt.Print("Ошибка! Повторите ввод:")
		fmt.Scan(&operation)
	}
	return operation
}

func getNums() ([]float64, error) {
	var numsStr string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&numsStr)
	numsStrSli := strings.Split(strings.TrimSpace(numsStr), ",")
	numsSli := []float64{}
	for _, str := range numsStrSli {
		num, error := strconv.ParseFloat(str, 32)
		if error != nil {
			return nil, error
		}
		numsSli = append(numsSli, num)
	}
	return numsSli, nil
}

func outputResult(operation string, numsSli *[]float64) {
	numsLen := len(*numsSli)

	switch operation {
	case "AVG":
		acc := 0.0
		for _, n := range *numsSli {
			acc += n
		}
		res := acc / float64(numsLen)
		fmt.Printf("Среднее ваших чисел = %.2f\n", res)
	case "SUM":
		acc := 0.0
		for _, n := range *numsSli {
			acc += n
		}
		fmt.Printf("Сумма ваших чисел = %.2f\n", acc)
	case "MED":
		res := 0.0
		slices.Sort(*numsSli)
		if numsLen%2 == 0 {
			halfLen := numsLen / 2
			res = ((*numsSli)[halfLen-1] + (*numsSli)[halfLen]) / 2
		} else {
			halfLen := float64(numsLen) / 2
			res = (*numsSli)[int(halfLen)]
		}
		fmt.Printf("Медиана ваших чисел = %.2f\n", res)
	}
}

func main() {
	fmt.Println("___ Калькулятор ___")
	operation := getOperation()

	numsSli, err := getNums()
	for err != nil {
		fmt.Println("Ошибка! Повторите ввод")
		numsSli, err = getNums()
	}

	outputResult(operation, &numsSli)
}
