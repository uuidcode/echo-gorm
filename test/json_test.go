package test

import (
	"fmt"
	"github.com/echo-gorm/util"
	"reflect"
	"testing"
)

type Org struct {
	Name string
}

type Sector struct {
	Name string
	Org  Org
}

type Team struct {
	Name   string
	Sector Sector
}

type Part struct {
	Name string
	Team Team
}

func look(item interface{}) {
	fmt.Println("typeOf", reflect.TypeOf(item))
	fmt.Println("kind", reflect.TypeOf(item).Kind())
	fmt.Println("json", util.ToJson(item))
}

func TestJson(t *testing.T) {
	part := new(Part)
	look(part)

	var part2 Part
	look(part2)

	part3 := Part{}
	look(part3)
}
