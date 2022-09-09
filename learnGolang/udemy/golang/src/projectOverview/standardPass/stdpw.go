package standardPass

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

var (
	ErrInvalidLength        = &ErrInvalidLengthType{}
	ErrMissingSmallLetter   = errors.New("Password is missing a small letter")
	ErrMissingCapitalLetter = errors.New("Password is missing a capital letter")
	ErrMissingDigit         = &ErrMissingDigitType{}
)

type PasswordError struct {
	msg string
	err error
}

type ErrInvalidLengthType struct {
	ActualLength int
	Min          int
	Max          int
}

type ErrMissingDigitType struct {
	Desc string
}

func (e ErrInvalidLengthType) Error() string {
	return "Password is invalid length"
}

func (e ErrMissingDigitType) Error() string {
	return "Password is missing a digital letter"
}

func (pe *PasswordError) Error() string {
	return pe.msg
}

func (pe *PasswordError) Unwrap() error {
	return pe.err
}

func (pe *PasswordError) wrappedBy(e error) {
	pe.msg = pe.msg + e.Error() + "\n"

	// e is wrapping pe.
	// pe.err = fmt.Errorf("%w", e)
	pe.err = &Unwrappable{wrapper: e, wrapped: pe.err}
}

type Unwrappable struct {
	wrapper error
	wrapped error
}

func (e *Unwrappable) As(target interface{}) bool {
	switch target.(type) {
	case *ErrInvalidLengthType:
		if src, ok := e.wrapper.(*ErrInvalidLengthType); ok {
			v := target.(*ErrInvalidLengthType)
			v.ActualLength = src.ActualLength
			v.Max = src.Max
			v.Min = src.Min
			return true
		}
	case *ErrMissingDigitType:
		if src, ok := e.wrapper.(*ErrMissingDigitType); ok {
			v := target.(*ErrMissingDigitType)
			v.Desc = src.Desc
			return true
		}
	}

	return false
}

func (e *Unwrappable) Error() string {
	return e.wrapper.Error() // error ชั้นนอกสุด
}

func (pe *PasswordError) Is(err error) bool {
	// check Error string
	return pe.err.Error() == err.Error()
}

func (e *Unwrappable) Unwrap() error {
	return e.wrapped // error ชั้นถัดไป
}

func (e *Unwrappable) Wrapper() error {
	return e.wrapper
}

func PasswordValidation(pw string) error {
	// check length
	pwError := &PasswordError{}
	if e := checkLength(pw); e != nil {
		pwError.wrappedBy(e) // error จะถูก wrap ด้วย e
	}

	// check contain small letter
	if e := containSmallLetter(pw); e != nil {
		pwError.wrappedBy(e)
	}

	// check contain capital letter
	if e := containCapitalLetter(pw); e != nil {
		pwError.wrappedBy(e)
	}

	// check contain digit
	if e := containDigit(pw); e != nil {
		pwError.wrappedBy(e)
	}

	if pwError.msg != "" {
		return pwError
	}

	return nil
}

func checkLength(pw string) error {
	pwLen := utf8.RuneCountInString(pw)
	if pwLen < 7 || pwLen > 16 {
		return &ErrInvalidLengthType{
			ActualLength: pwLen,
			Min:          7,
			Max:          16,
		}

	}
	return nil
}

func regexHelper(pw, pattern, msg string) error {
	re := regexp.MustCompile(pattern)
	result := re.FindString(pw)
	if result == "" {
		return fmt.Errorf(msg)
	}
	return nil
}

func containSmallLetter(pw string) error {
	// [a-z]
	err := regexHelper(pw, `[a-z]`, "Password must contain small letter")
	if err != nil {
		return ErrMissingSmallLetter
	}
	return nil
}

func containCapitalLetter(pw string) error {
	// `[A-Z]`
	err := regexHelper(pw, `[A-Z]`, "Password must contain capital letter")
	if err != nil {
		return ErrMissingCapitalLetter
	}
	return nil
}

func containDigit(pw string) error {
	err := regexHelper(pw, `[0-9]`, "Password must contain digit")
	if err != nil {
		return &ErrMissingDigitType{Desc: "Password must contain digit"}
	}
	return nil
}
