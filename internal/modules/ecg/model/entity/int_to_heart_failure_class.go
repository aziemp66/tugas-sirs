package ecg_entity

import (
	"errors"
	pkg_error "tugas-sirs/pkg/error"
)

func IntToHeartFailureClass(num int) (HeartFailureClass,error) {
	switch num {
	case 1:
		return classA,nil
	case 2:
		return classB,nil
	case 3:
		return classC,nil
	case 4:
		return classD,nil
	case 5:
		return classE,nil
	case 6:
		return classF,nil
	case 7:
		return classG,nil
	default :
		return "",pkg_error.NewBadRequest(
			errors.New("integer is not a valid heart failure class category"),
			"Heart Class Category Is Invalid",
		)
	}
}