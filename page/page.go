package page

const limit = int64(10)
const itemCount = int64(10)

type Page struct {
	Previous bool
	Next     bool
	List     []int64
	Total    int64
}

func New(p int64, total int64) *Page {
	item := Page{}
	item.Total = total
	totalPage := total/limit + 1

	if p > itemCount {
		item.Previous = true
	} else {
		item.Previous = false
	}

	item.Previous = p > itemCount
	item.Next = totalPage > ((p+itemCount)/itemCount)*itemCount

	var list []int64

	for i := int64(1); i <= itemCount; i++ {
		currentP := p + i

		if currentP <= totalPage {
			list = append(list, currentP)
		}
	}

	item.List = list

	return &item
}
