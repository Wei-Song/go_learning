package error

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("n should not be less than 2")
var LargeThanHundredError = errors.New("n should be not large than 100")

func GetFibonacci(n int) ([]int,error) {

	if n < 0 {
		return nil, LessThanTwoError
	}

	if  n > 100 {
		return nil, LargeThanHundredError
	}

	fibList := []int{1,1}

	for i := 2;i<n;i++{
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if fibonacci, err := GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(fibonacci)
	}

}
