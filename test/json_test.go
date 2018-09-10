package test

import (
	"encoding/json"
	"fmt"
	"github.com/echo-gorm/util"
	"reflect"
	"testing"
)

type Dummy struct {
	Price int
}

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
	Name        string
	Team        Team
	AnotherPart *Part `json:"anotherPart,omitempty"`
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

func TestStringToInt(t *testing.T) {
	text := `{"price":1000}`

	dummy := new(Dummy)
	err := json.Unmarshal([]byte(text), dummy)
	util.CheckErr(err)

	fmt.Println(util.ToJson(dummy))
}
