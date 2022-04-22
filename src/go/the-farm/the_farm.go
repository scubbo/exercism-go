package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	message string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	fodder, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction {
		if fodder < 0 {
			return 0, errors.New("negative fodder")
		}
		// a float number of cows? Madness. Why can't we just floor-divide?
		return fodder * 2 / float64(cows), nil
	}
	if err != nil {
		return 0, err
	}
	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, &SillyNephewError{message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
	}

	// Technically, the requirements don't say to return anything if there is no error.
	return fodder / float64(cows), nil
}
