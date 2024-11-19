//go:build go1.18

package clause

import (
	"github.com/hopeio/utils/types/param"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PageSort param.PageSort

func (req *PageSort) Clause() []clause.Expression {
	if req.PageNo == 0 && req.PageSize == 0 {
		return nil
	}
	if req.Sort == nil || req.Sort.SortField == "" {
		return []clause.Expression{PageExpr(req.PageNo, req.PageSize)}
	}

	return []clause.Expression{SortExpr(req.SortField, req.SortType), PageExpr(req.PageNo, req.PageSize)}
}

func FindByList[T any, O param.Ordered](db *gorm.DB, req *param.List[O]) ([]T, error) {
	var models []T
	clauses := ListClause(req)
	if len(clauses) > 0 {
		db = db.Clauses(clauses...)
	}
	err := db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func ListClause[O param.Ordered](req *param.List[O]) []clause.Expression {
	return (*List[O])(req).Clause()
}

type List[T param.Ordered] param.List[T]

func (req *List[O]) Clause() []clause.Expression {
	pqc := (*PageSort)(&req.PageSort).Clause()
	rqc := (*Range[O])(req.Range).Clause()
	if pqc != nil && rqc != nil {
		return append((*PageSort)(&req.PageSort).Clause(), (*Range[O])(req.Range).Clause())
	}
	if rqc == nil {
		return pqc
	}
	if rqc != nil {
		return []clause.Expression{rqc}
	}
	return nil
}

func PageExpr(pageNo, pageSize int) clause.Limit {
	if pageSize == 0 {
		pageSize = 100
	}
	if pageNo > 1 {
		return clause.Limit{Offset: (pageNo - 1) * pageSize, Limit: &pageSize}
	}
	return clause.Limit{Limit: &pageSize}
}
