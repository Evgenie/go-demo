package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

var operationsMap = map[string]func(*[]float64){
	"AVG": outputAvg,
	"SUM": outputSum,
	"MED": outputMed,
}

var operations = slices.Collect(maps.Keys(operationsMap))

func outputAvg(numsSli *[]float64) {
	acc := 0.0
	for _, n := range *numsSli {
		acc += n
	}
	res := acc / float64(len(*numsSli))
	fmt.Printf("Среднее ваших чисел = %.2f\n", res)
}
func outputSum(numsSli *[]float64) {
	acc := 0.0
	for _, n := range *numsSli {
		acc += n
	}
	fmt.Printf("Сумма ваших чисел = %.2f\n", acc)
}
func outputMed(numsSli *[]float64) {
	res := 0.0
	slices.Sort(*numsSli)
	numsLen := len(*numsSli)
	if numsLen%2 == 0 {
		halfLen := numsLen / 2
		res = ((*numsSli)[halfLen-1] + (*numsSli)[halfLen]) / 2
	} else {
		halfLen := float64(numsLen) / 2
		res = (*numsSli)[int(halfLen)]
	}
	fmt.Printf("Медиана ваших чисел = %.2f\n", res)
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
	action := operationsMap[operation]
	if action != nil {
		action(numsSli)
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
