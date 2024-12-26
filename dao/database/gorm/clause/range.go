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
	if req.RangeType == 0 {
		return NewWhereClause(req.RangeField, dbi.Between, req.RangeBegin, req.RangeEnd)
	}
	if req.RangeType.HasBegin() && req.RangeType.HasEnd() {
		if req.RangeType.ContainsBegin() && req.RangeType.ContainsEnd() {
			return NewWhereClause(req.RangeField, dbi.Between, req.RangeBegin, req.RangeEnd)
		} else {
			leftOp, rightOp := dbi.Greater, dbi.Less
			if req.RangeType.ContainsBegin() {
				leftOp = dbi.GreaterOrEqual
			}
			if req.RangeType.ContainsEnd() {
				leftOp = dbi.LessOrEqual
			}
			return clause.Where{Exprs: []clause.Expression{NewWhereClause(req.RangeField, leftOp, req.RangeBegin), NewWhereClause(req.RangeField, rightOp, req.RangeEnd)}}
		}
	}

	if req.RangeType.HasBegin() {
		operation := dbi.Greater
		if req.RangeType.ContainsEnd() {
			operation = dbi.GreaterOrEqual
		}
		return NewWhereClause(req.RangeField, operation, req.RangeBegin)
	}
	if req.RangeType.HasEnd() {
		operation := dbi.Less
		if req.RangeType.ContainsEnd() {
			operation = dbi.LessOrEqual
		}
		return NewWhereClause(req.RangeField, operation, req.RangeEnd)
	}

	return nil
}
