package test

import (
	"fmt"
	"github.com/echo-gorm/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {
	var m map[string]string
	assert.Nil(t, m)
	fmt.Println(util.ToJson(m))

	m1 := map[string]string{}
	assert.NotNil(t, m1)
	fmt.Println(util.ToJson(m1))

	m2 := map[string]string{
		"1": "1",
	}
	assert.NotNil(t, m2)
	fmt.Println(util.ToJson(m2))

	m3 := make(map[string]string)
	assert.NotNil(t, m3)
	fmt.Println(util.ToJson(m3))
}

func TestChange(t *testing.T) {
	m := map[string]string{
		"1": "1",
	}

	change(m)

	fmt.Println(util.ToJson(m))
}

func change(m map[string]string) {
	m["1"] = "2"
}
