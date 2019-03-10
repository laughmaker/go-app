package app

import (
	"math"

	"github.com/gin-gonic/gin"
)

type Page struct {
	PageIndex int `json:"pageIndex"` // 分页索引
	PageSize  int `json:"pageSize"`  // 分页大小
	PageCount int `json:"pageCount"` // 分页数量
	TotalRows int `json:"totalRows"` // 总行数
}

func NewPage(c *gin.Context) Page {
	p := Page{}
	p.PageIndex = c.GetInt(c.PostForm("pageIndex"))
	p.PageSize = c.GetInt(c.PostForm("pageSize"))

	return p
}

func (p *Page) Limit() int {
	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}

	return p.PageSize
}

func (p *Page) Offset() int {
	return p.PageIndex * p.Limit()
}

func (p *Page) Update() {
	p.PageCount = int(math.Ceil(float64(p.TotalRows) / float64(p.Limit())))
	p.PageSize = p.Limit()
}
