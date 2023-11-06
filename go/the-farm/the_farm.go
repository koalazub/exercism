package thefarm

import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, qty int) (float64, error) {

	fa, err := f.FodderAmount(qty)
	if err != nil {
		return 0, err
	}

	ff, err := f.FatteningFactor()
	if err != nil {
		return 0, err
	}

	fdrPerCow := fa / float64(qty)

	return ff * fdrPerCow, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function

func ValidateInputAndDivideFood(f FodderCalculator, qty int) (float64, error) {
	if qty <= 0 {
		return 0, fmt.Errorf("invalid number of cows")
	}

	// DivideFood returns error as well and so you don't need nil here
	return DivideFood(f, qty)
}

type InvalidCowsError error

type Error interface {
	CowError(qty int, msg string)
}

func CowError(qty int, msg string) InvalidCowsError {
	return InvalidCowsError(fmt.Errorf("%d cows are invalid: %v", qty, msg))
}

func ValidateNumberOfCows(qty int) InvalidCowsError {
	if qty < 0 {
		return CowError(qty, "there are no negative cows")
	}
	if qty == 0 {
		return CowError(qty, "no cows don't need food")
	}

	return nil
}

// TODO: define the 'ValidateNumberOfCows' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
