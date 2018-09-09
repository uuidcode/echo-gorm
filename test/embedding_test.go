package test

import (
	"fmt"
	"github.com/echo-gorm/util"
	"testing"
)

type Reader struct {
	In string
}

type Writer struct {
	Out string
}

type File struct {
	Reader
	Writer
}

func TestName(t *testing.T) {
	var file File
	file.In = "In"
	file.Out = "Out"

	fmt.Println(util.ToJson(file))
}
