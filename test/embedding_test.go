package test

import (
	"fmt"
	"github.com/echo-gorm/util"
	"testing"
)

type Reader struct {
	In string
}

func (Reader) Read() {
	fmt.Println("Read in Reader")
}

type Writer struct {
	Out string
}

type File struct {
	Reader
	Writer
}

func (File) Read() {
	fmt.Println("Read in File")
}

func TestEmbedding(t *testing.T) {
	var file File
	file.In = "In"
	file.Out = "Out"
	file.Read()

	fmt.Println(util.ToJson(file))
}
