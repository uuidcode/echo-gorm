package test

import (
	"fmt"
	"testing"
)

func TestUnit(t *testing.T) {
	defer func() {
		fmt.Println("defer")

		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	intValue := -1
	value := uint(intValue)
	fmt.Println(value)
}
