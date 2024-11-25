//go:build go1.18

/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package clause

import (
	dbi "github.com/hopeio/utils/dao/database"
	"github.com/hopeio/utils/types/param"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewWhereClause(field string, op dbi.Operation, args ...interface{}) clause.Expression {
	switch op {
	case dbi.Equal:
		return clause.Eq{
			Column: field,
			Value:  args[0],
		}
	case dbi.In:
		return clause.IN{
			Column: field,
			Values: args,
		}
	case dbi.Between:
		return clause.Expr{
			SQL:  field + " BETWEEN ? AND ?",
			Vars: args,
		}
	case dbi.Greater:
		return clause.Gt{
			Column: field,
			Value:  args[0],
		}
	case dbi.Less:
		return clause.Lt{
			Column: field,
			Value:  args[0],
		}
	case dbi.LIKE:
		return clause.Like{
			Column: field,
			Value:  args[0],
		}
	case dbi.GreaterOrEqual:
		return clause.Gte{
			Column: field,
			Value:  args[0],
		}
	case dbi.LessOrEqual:
		return clause.Lte{
			Column: field,
			Value:  args[0],
		}
	case dbi.NotIn:
		return clause.NotConditions{Exprs: []clause.Expression{clause.IN{
			Column: field,
			Values: args,
		}}}
	case dbi.NotEqual:
		return clause.Neq{
			Column: field,
			Value:  args[0],
		}
	case dbi.IsNull:
		return clause.Expr{
			SQL:  field + " IS NULL",
			Vars: nil,
		}
	case dbi.IsNotNull:
		return clause.Expr{
			SQL:  field + " IS NOT NULL",
			Vars: nil,
		}
	}
	return clause.Expr{
		SQL:  field,
		Vars: args,
	}
}

func DateBetween(column, dateStart, dateEnd string) clause.Expression {
	return NewWhereClause(column, dbi.Between, dateStart, dateEnd)
}

func SortExpr(column string, typ param.SortType) clause.Expression {
	var desc bool
	if typ == param.SortTypeDesc {
		desc = true
	}
	return clause.OrderBy{Columns: []clause.OrderByColumn{{Column: clause.Column{Name: column, Raw: true}, Desc: desc}}}
}

func TableName(tx *gorm.DB, name string) *gorm.DB {
	tx.Statement.TableExpr = &clause.Expr{SQL: tx.Statement.Quote(name)}
	tx.Statement.Table = name
	return tx
}

type Expression dbi.FilterExpr

func (e *Expression) Clause() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(e.Field+(*dbi.FilterExpr)(e).Operation.SQL(), e.Value...)
	}
}

func ByValidEqual[T comparable](column string, v T) clause.Expression {
	var zero T
	if v != zero {
		return clause.Eq{Column: column, Value: v}
	}
	return nil
}
