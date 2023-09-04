package e2

import (
	"errors"
	"fmt"
)

func main() {
	notes := []float64{5.0, 4.5, 3.0, 5}
	mean, err := CalculateMean(notes...)
	fmt.Println("Mean is", mean, err)
}

func CalculateMean(notes ...float64) (mean float64, err error) {
	result := 0.0
	for _, value := range notes {
		if value < 0 {
			err = errors.New("It can't be negative")
		}
	}
	if err == nil {
		for _, value := range notes {
			result += value
		}
		mean = result / float64(len(notes))
		err = errors.New("")
	}
	return
}
