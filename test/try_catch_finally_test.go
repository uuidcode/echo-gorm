package test

/*
 * https://dzone.com/articles/try-and-catch-in-golang
 */
import (
	"fmt"
	"testing"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (block Block) Do() {
	if block.Finally != nil {
		defer block.Finally()
	}

	if block.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				block.Catch(r)
			}
		}()
	}

	block.Try()
}

func TestTryCatchFinally(t *testing.T) {
	fmt.Println("We started")

	Block{
		Try: func() {
			fmt.Println("Try")
			Throw("Error")
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Do()

	fmt.Println("We went on")
}
