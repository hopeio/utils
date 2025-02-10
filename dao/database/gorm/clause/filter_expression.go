package clause

import (
	dbi "github.com/hopeio/utils/dao/database/sql"
	"gorm.io/gorm/clause"
	"strings"
)

type FilterExpr dbi.FilterExpr

func (f *FilterExpr) Clause() clause.Expression {
	f.Field = strings.TrimSpace(f.Field)

	return NewCondition(f.Field, f.Operation, f.Value...)
}

type FilterExprs dbi.FilterExprs

func (f FilterExprs) Clause() clause.Expression {
	var exprs []clause.Expression
	for _, filter := range f {
		filter.Field = strings.TrimSpace(filter.Field)

		filterExpr := (FilterExpr)(filter)
		expr := filterExpr.Clause()
		if expr != nil {
			exprs = append(exprs, expr)
		}
	}
	if len(exprs) > 0 {
		return clause.AndConditions{Exprs: exprs}
	}
	return nil
}
