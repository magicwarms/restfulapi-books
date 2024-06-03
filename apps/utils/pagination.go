package utils

import (
	"gorm.io/gorm"
)

type paginate struct {
	perPage int
	page    int
}

func NewPaginate(page, perPage int) *paginate {
	return &paginate{perPage: perPage, page: page}
}

func (p *paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * int(p.perPage)

	return db.Offset(int(offset)).
		Limit(p.perPage)
}
