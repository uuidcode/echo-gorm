package test

import (
	"github.com/echo-gorm/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Student struct {
	Name string
}

func TestStudent1(t *testing.T) {
	student := Student{
		Name: "Hello",
	}

	changeStudentByValue(student)
	util.PrintTypeAndPointer(student)
	assert.Equal(t, "Hello", student.Name)
}

func TestStudent2(t *testing.T) {
	student := new(Student)
	student.Name = "Hello"
	changeStudentByPointer(student)
	util.PrintTypeAndPointer(student)
	assert.Equal(t, "World", student.Name)
}

func changeStudentByValue(student Student) {
	util.PrintTypeAndPointer(student)
	student.Name = "World"
}

func changeStudentByPointer(student *Student) {
	util.PrintTypeAndPointer(student)
	student.Name = "World"
}
