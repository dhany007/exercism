package thefarm

import (
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var (
		errDivideZero error   = fmt.Errorf("division by zero")
		errNegative   error   = fmt.Errorf("negative fodder")
		zeroFodder    float64 = 0.0
	)

	fodder, err := weightFodder.FodderAmount()

	if err != nil {
		switch err {
		case ErrScaleMalfunction:
			fodder *= 2
		default:
			return zeroFodder, err
		}
	}

	if fodder < 0 {
		return zeroFodder, errNegative
	}

	if cows == 0 {
		return zeroFodder, errDivideZero
	}

	if cows < 0 {
		return zeroFodder, &SillyNephewError{cows}
	}

	return fodder / float64(cows), nil
}
