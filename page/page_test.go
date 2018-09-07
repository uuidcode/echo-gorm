package page

import (
	"fmt"
	"github.com/echo-gorm/util"
	"testing"
)

func TestMath(t *testing.T) {
	page := New(0, 9)
	fmt.Println(util.ToJson(page))
}
