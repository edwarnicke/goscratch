package errorshadow_test

import (
	"errors"
	"fmt"
	"testing"
	. "github.com/onsi/gomega"
)

const (
	testErrorString = "This is a test error"
)


func returnStringErr() (s string, err error) {
	s = "This is a test string"
	err = errors.New(testErrorString)
	return s,err
}

func returnErrCompouncRedeclare() (err error) {
	// In this case, err is not redeclared
	s,err := returnStringErr()
	fmt.Println("returnErrCompouncRedeclare: %s",s)
	return err
}
func returnErrCompouncRedeclareInIf01() (err error) {
	// In this case, err is redeclared, because its in the if scope
	if s,err := returnStringErr();err != nil {
		fmt.Println("returnErrCompouncRedeclare: %s",s)
	}
	return err
}

func returnErrCompouncRedeclareInIf02() (err error) {
	if true {
		// In this case err is redeclared because its in the if scope
		s,err := returnStringErr()
		fmt.Println("returnErrCompouncRedeclare err:",err)
		fmt.Println("returnErrCompouncRedeclare: s:",s)
	}
	return err
}


func TestErrorShadowing(t *testing.T) {
	RegisterTestingT(t)
	err := returnErrCompouncRedeclare()
	Expect(err).ShouldNot(BeNil())
	Expect(err.Error()).Should(Equal(testErrorString))
}

func TestErrorShadowingInIf01(t *testing.T) {
	RegisterTestingT(t)
	err := returnErrCompouncRedeclareInIf01()
	Expect(err).Should(BeNil())
}

func TestErrorShadowingInIf02(t *testing.T) {
	RegisterTestingT(t)
	err := returnErrCompouncRedeclareInIf01()
	Expect(err).Should(BeNil())
}
