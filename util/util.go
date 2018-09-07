package util

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"regexp"
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

func GetPage(c echo.Context) int64 {
	pValue := c.QueryParam("p")
	p, err := ParseInt(pValue)

	if err != nil {
		p = 1
	}

	if p == 0 {
		p = 1
	}

	return p
}

func GetUrl(req *http.Request) string {
	path := req.URL.Path
	query := req.URL.RawQuery
	return fmt.Sprintf("%s?%s", path, query)
}

func GetUrlAndRemovePageParam(req *http.Request) string {
	path := req.URL.Path
	query := req.URL.RawQuery

	reg, err := regexp.Compile("&?p=[0-9]*")
	CheckErr(err)

	query = reg.ReplaceAllString(query, "")

	url := path + "?"

	if query != "" {
		url += query

		if !strings.HasSuffix(url, "&") {
			url += "&"
		}
	}

	return url
}
