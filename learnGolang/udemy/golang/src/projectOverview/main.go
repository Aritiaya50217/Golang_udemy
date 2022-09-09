package main

import (
	stdpw "ProjectOverview/standardPass"
	"errors"
	"fmt"
)

type Unwrappable struct {
	msg     string
	wrapped error
}

func (e Unwrappable) Error() string {
	return e.msg
}

func (e Unwrappable) Unwrap() error {
	return e.wrapped
}

func main() {
	if err := stdpw.PasswordValidation("#123oilS"); err != nil {
		fmt.Printf("%v\n", *(err.(*stdpw.PasswordError)))
		// handling error when invalid Length
		if errors.Is(err, stdpw.ErrInvalidLength) {
			// Client code can handle their own business login
			fmt.Println("custom : Invalid Length")

			errInvalidLength := &stdpw.ErrInvalidLengthType{}
			if errors.As(err, errInvalidLength) {
				fmt.Println("\t--->", errInvalidLength.ActualLength)
			}
		}

		if errors.Is(err, stdpw.ErrMissingSmallLetter) {
			fmt.Println("custom : MissingSmallLetter")
		}

		if errors.Is(err, stdpw.ErrMissingCapitalLetter) {
			fmt.Println("custom : MissingCapitalLetter")
		}
		if errors.Is(err, stdpw.ErrMissingDigit) {
			fmt.Println("custom : MissingDigit")
			errMissingDigit := &stdpw.ErrMissingDigitType{}
			if errors.As(err, errMissingDigit) {
				fmt.Println("\t--->", errMissingDigit.Desc)
			}
		}
	}

	fmt.Println("====== Wrap ======")

	errC := fmt.Errorf("I am error C")
	// %w => wrap
	errA := fmt.Errorf("I am error A %w", errC) // errA wrap errC

	errorIs := errors.Is(errA, errC)
	fmt.Println("Error is :", errorIs)
	fmt.Println("Done")

	fmt.Println("====== Unwrap ======")

	errC2 := fmt.Errorf("I am error C2 ")
	errA2 := Unwrappable{msg: "I am error A2", wrapped: errC2}
	errorIs2 := errors.Is(errA2, errC2)
	fmt.Println("Error is :", errorIs2)

}

/*
	Password validation

			[ 7-16 charactor ]       [ Contain Small Charater ]			[ Contain Capital Character ]			[Contain digit ]

[123a] 		====> checkLength		 ====> checkMissingSmallLetter		====> checkMissingCapitalLetter		 	====> checkMissingDigit

result : ErrMissingDigit


*/
