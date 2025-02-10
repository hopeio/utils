//go:build go1.18

/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package clause

import (
	dbi "github.com/hopeio/utils/dao/database/sql"
	"github.com/hopeio/utils/types/param"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewWhereClause(field string, op dbi.Operation, args ...any) clause.Expression {
	if field == "" {
		return nil
	}
	switch op {
	case dbi.Equal:
		if len(args) == 0 {
			return nil
		}
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
		if len(args) != 2 {
			return nil
		}
		return Between{
			Column: field,
			Begin:  args[0],
			End:    args[1],
		}
	case dbi.Greater:
		if len(args) == 0 {
			return nil
		}
		return clause.Gt{
			Column: field,
			Value:  args[0],
		}
	case dbi.Less:
		if len(args) == 0 {
			return nil
		}
		return clause.Lt{
			Column: field,
			Value:  args[0],
		}
	case dbi.LIKE:
		if len(args) == 0 {
			return nil
		}
		return clause.Like{
			Column: field,
			Value:  args[0],
		}
	case dbi.GreaterOrEqual:
		if len(args) == 0 {
			return nil
		}
		return clause.Gte{
			Column: field,
			Value:  args[0],
		}
	case dbi.LessOrEqual:
		if len(args) == 0 {
			return nil
		}
		return clause.Lte{
			Column: field,
			Value:  args[0],
		}
	case dbi.NotIn:
		return Not{Expr: clause.IN{
			Column: field,
			Values: args,
		}}
	case dbi.NotEqual:
		if len(args) == 0 {
			return nil
		}
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

func ByPrimaryKey(v any) clause.Expression {
	return clause.Eq{
		Column: clause.PrimaryColumn,
		Value:  v,
	}
}

type Between struct {
	Column     any
	Begin, End any
}

func (gt Between) Build(builder clause.Builder) {
	builder.WriteQuoted(gt.Column)
	builder.WriteString(" BETWEEN ")
	builder.AddVar(builder, gt.Begin)
	builder.WriteString(" AND ")
	builder.AddVar(builder, gt.End)
}

func (gt Between) NegationBuild(builder clause.Builder) {
	builder.WriteQuoted(gt.Column)
	builder.WriteString(" < ")
	builder.AddVar(builder, gt.Begin)
	builder.WriteString(" OR ")
	builder.WriteQuoted(gt.Column)
	builder.WriteString(" > ")
	builder.AddVar(builder, gt.End)
}

type Not struct {
	Expr clause.NegationExpressionBuilder
}

func (n Not) Build(builder clause.Builder) {
	n.Expr.NegationBuild(builder)
}
