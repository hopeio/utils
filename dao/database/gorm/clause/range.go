/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package clause

import (
	dbi "github.com/hopeio/utils/dao/database"
	"github.com/hopeio/utils/types/param"
	"gorm.io/gorm/clause"
)

type Range[T param.Ordered] param.Range[T]

func (req *Range[T]) Clause() clause.Expression {
	if req == nil || req.RangeField == "" {
		return nil
	}

	var zero T
	operation := dbi.Between
	if req.RangeEnd == zero && req.RangeBegin != zero {
		operation = dbi.Greater
		if req.Include {
			operation = dbi.GreaterOrEqual
		}
		return NewWhereClause(req.RangeField, operation, req.RangeBegin)
	}
	if req.RangeBegin == zero && req.RangeEnd != zero {
		operation = dbi.Less
		if req.Include {
			operation = dbi.LessOrEqual
		}
		return NewWhereClause(req.RangeField, operation, req.RangeBegin)
	}
	if req.RangeBegin != zero && req.RangeEnd != zero {
		if req.Include {
			return NewWhereClause(req.RangeField, operation, req.RangeBegin, req.RangeEnd)
		} else {
			return clause.Where{Exprs: []clause.Expression{NewWhereClause(req.RangeField, dbi.Greater, req.RangeBegin), NewWhereClause(req.RangeField, dbi.Less, req.RangeBegin)}}
		}
	}
	return nil
}
