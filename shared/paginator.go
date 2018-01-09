package shared

import (
	"net/http"
	"strconv"
)

type Paginator struct {
	limit  int
	offset int
}

func NewPaginatorFromRequest(r *http.Request) *Paginator {
	page, pageErr := strconv.Atoi(r.URL.Query().Get("page"))
	limit, limitError := strconv.Atoi(r.URL.Query().Get("limit"))

	if pageErr != nil || page < -1 {
		page = 1
	}

	if limitError != nil || limit > 50 || limit < 1 {
		limit = 10
	}

	offset := (page * limit) - 1

	if page == 1 {
		offset = 0
	}

	return &Paginator{limit, offset}
}

func (paginator Paginator) Limit() int {
	return paginator.limit
}

func (paginator Paginator) Offset() int {
	return paginator.offset
}
