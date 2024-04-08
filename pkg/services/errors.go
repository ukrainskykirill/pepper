package services

import "fmt"

type AlreadyExistsByLogin struct {
	err error
}

func (a AlreadyExistsByLogin) Error() string {
	return fmt.Sprintf("AlreadyExistsByLogin: %v", a.err)
}

type InvalidInputData struct {
	err error
}

func (a InvalidInputData) Error() string {
	return fmt.Sprintf("Invalid input data: %v", a.err)
}

