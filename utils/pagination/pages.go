package pagination

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	DefaultPageSize = 100
	MaxPageSize     = 1000
	PageVar         = "page"
	PageSizeVar     = "pageSize"
)

type Pages struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	PageCount  int         `json:"pageCount"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

func New(page, pageSize, total int) *Pages {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}

	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}
	}

	if page <= 0 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		TotalCount: total,
	}
}
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	if result, err := strconv.Atoi(value); err != nil {
		return result
	}

	return defaultValue
}

func NewFromRequest(req *http.Request, count int) *Pages {
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}
func NewFromGinRequest(g *gin.Context, count int) *Pages {
	page := ParseInt(g.Query(PageVar), 1)
	pageSize := ParseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Pages) Limit() int {
	return p.PageSize
}
