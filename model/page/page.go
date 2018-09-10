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
	TotalItem    int64
	ItemLimit    int64
	ItemOffset   int64
	PageList     []int64
	CurrentPage  int64
	PageLimit    int64
}

func NewWithContext(c echo.Context, total int64) *Page {
	p := util.GetPage(c)
	return NewWithLimitAndItemCount(p, total, 10, 10)
}

func New(p int64, total int64, itemCount int64) *Page {
	return NewWithLimitAndItemCount(p, total, 10, itemCount)
}

func NewWithLimitAndItemCount(currentPage int64, totalItem int64, itemLimit int64, pageLimit int64) *Page {
	if currentPage < 1 {
		currentPage = 1
	}

	item := Page{}
	item.ItemLimit = itemLimit
	item.PageLimit = pageLimit
	item.CurrentPage = currentPage
	item.ItemOffset = (currentPage - 1) * itemLimit
	item.TotalItem = totalItem
	totalPage := int64(math.Ceil(float64(totalItem) / float64(itemLimit)))

	var list []int64

	startPage := ((currentPage - 1) / pageLimit) * pageLimit

	for i := int64(1); i <= pageLimit; i++ {
		currentP := startPage + i

		if currentP <= totalPage {
			list = append(list, currentP)
		}
	}

	item.Previous = currentPage > pageLimit
	item.Next = totalPage > ((currentPage+pageLimit)/pageLimit)*pageLimit

	if item.Previous {
		item.PreviousPage = list[0] - pageLimit
	}

	if item.Next {
		item.NextPage = list[len(list)-1] + 1
	}

	item.PageList = list

	return &item
}
