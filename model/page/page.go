package page

import (
	"github.com/echo-gorm/util"
	"github.com/labstack/echo"
	"math"
)

type Page struct {
	Previous     bool
	Next         bool
	PreviousPage int64
	NextPage     int64
	List         []int64
	Total        int64
	CurrentPage  int64
	Limit        int64
	ItemCount    int64
	Offset       int64
}

func NewWithContext(c echo.Context, total int64) *Page {
	p := util.GetPage(c)
	return NewWithLimitAndItemCount(p, total, 10, 10)
}

func New(p int64, total int64, itemCount int64) *Page {
	return NewWithLimitAndItemCount(p, total, 10, itemCount)
}

func NewWithLimitAndItemCount(p int64, total int64, limit int64, itemCount int64) *Page {
	if p < 1 {
		p = 1
	}

	item := Page{}
	item.Limit = limit
	item.ItemCount = itemCount
	item.CurrentPage = p
	item.Offset = (p - 1) * limit
	item.Total = total
	totalPage := int64(math.Ceil(float64(total) / float64(limit)))

	if p > itemCount {
		item.Previous = true
	} else {
		item.Previous = false
	}

	item.Previous = p > itemCount
	item.Next = totalPage > ((p+itemCount)/itemCount)*itemCount

	list := []int64{}

	startPage := ((p - 1) / itemCount) * itemCount

	for i := int64(1); i <= itemCount; i++ {
		currentP := startPage + i

		if currentP <= totalPage {
			list = append(list, currentP)
		}
	}

	if item.Previous {
		item.PreviousPage = list[0] - itemCount
	}

	if item.Next {
		item.NextPage = list[len(list)-1] + 1
	}

	item.List = list

	return &item
}
