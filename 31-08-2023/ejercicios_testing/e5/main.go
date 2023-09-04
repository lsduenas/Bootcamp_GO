package e5

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {

	dogFunc, err := Animal(dog)
	if err == nil {
		animalsAmount := dogFunc(6000)
		fmt.Println("Dogs amount", animalsAmount)
	}

	catFunc, err := Animal(cat)
	if err == nil {
		animalsAmount := catFunc(5000)
		fmt.Println("Cats amount", animalsAmount)
	}

	hamsterFunc, err := Animal(hamster)
	if err == nil {
		animalAmount := hamsterFunc(23345)
		fmt.Println("Hamsters amount", animalAmount)
	}

	tarantulaFunc, err := Animal(tarantula)
	if err == nil {
		animalAmount := tarantulaFunc(22000)
		fmt.Println("Tarantulas amount", animalAmount)
	}

}

func Animal(animal string) (func(animal float64) float64, error) {
	switch animal {
	case dog:
		{
			return AnimalDog, nil
		}
	case cat:
		{
			return AnimalCat, nil
		}
	case hamster:
		{
			return AnimalHamster, nil
		}
	case tarantula:
		{
			return AnimalTarantula, nil
		}
	default:
		err := errors.New("Not valid animal")
		return nil, err
	}
}

func AnimalDog(amount float64) float64 {
	return amount / 10.0
}

func AnimalCat(amount float64) float64 {
	return amount / 5.0
}

func AnimalHamster(amount float64) float64 {
	return amount / 250.0
}

func AnimalTarantula(amount float64) float64 {
	return amount / 150.0
}
