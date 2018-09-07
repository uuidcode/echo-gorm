package util

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateUuid() string {
	value := uuid.NewV4()
	uuid := value.String()
	uuid = strings.Replace(uuid, "-", "", -1)
	return uuid
}

func DeleteItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func ToJson(object interface{}) string {
	bytes, err := json.MarshalIndent(object, "", "    ")
	CheckErr(err)

	return string(bytes)
}

func ParseInt(value string) (i int64, err error) {
	return strconv.ParseInt(value, 10, 64)
}

func GetOffsetAndLimit(c echo.Context) (offset int64, limit int64) {
	p := GetPage(c)

	offset = p * 10
	limit = 10
	return
}

func GetPage(c echo.Context) int64 {
	pValue := c.QueryParam("p")
	p, err := ParseInt(pValue)

	if err != nil {
		p = 0
	}

	return p
}
