package thefarm

import (
	"errors"
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

// TODO: define the 'DivideFood' function
func DivideFood(fodderCalculator FodderCalculator, cowNumber int) (float64, error) {
	fodderAmount, err := fodderCalculator.FodderAmount(cowNumber)
	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, errors.New("something went wrong")
	}
	return fodderAmount * fatteningFactor / float64(cowNumber), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cowNumber int) (float64, error) {
	if cowNumber > 0 {
		return DivideFood(fodderCalculator, cowNumber)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cowNumber int
	message   string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cowNumber, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function

func ValidateNumberOfCows(cowNumber int) error {
	if cowNumber < 0 {
		return &InvalidCowsError{
			cowNumber: cowNumber,
			message:   "there are no negative cows",
		}
	}
	if cowNumber == 0 {
		return &InvalidCowsError{
			cowNumber: cowNumber,
			message:   "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
