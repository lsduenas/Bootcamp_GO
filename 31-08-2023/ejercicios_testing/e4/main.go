package e4

import (
	"errors"
	"fmt"
)

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func main() {
	minFunc, err := operation(Minimum)
	if err == nil {
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Min", minValue)
	}

	averageFunc, err := operation(Average)
	if err == nil {
		averageValue := averageFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Average", averageValue)
	}

	maxFunc, err := operation(Maximum)
	if err == nil {
		maxValue := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Println("Max", maxValue)
	}

}

func operation(operation string) (func(values ...float64) float64, error) {
	switch operation {
	case Minimum:
		{
			return OpMin, nil
		}
	case Average:
		{
			return OpAverage, nil
		}
	case Maximum:
		{
			return OpMax, nil
		}
	default:
		err := errors.New("Not valid operation")
		return nil, err
	}

}

func OpMin(values ...float64) (min float64) {
	min = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return
}

func OpAverage(values ...float64) (result float64) {
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	result = sum / float64(len(values))
	fmt.Println("olaaaaa", result)
	return
}

func OpMax(values ...float64) (max float64) {
	max = values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return
}
