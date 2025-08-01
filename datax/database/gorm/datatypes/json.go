/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package datatypes

import (
	"context"
	"database/sql/driver"
	dbi "github.com/hopeio/gox/datax/database"
	"github.com/hopeio/gox/datax/database/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Json[T any] datatypes.Json[T]

func (*Json[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case dbi.Sqlite, dbi.Mysql:
		return "json"
	case dbi.Postgres:
		return "jsonb"
	}
	return ""
}

func (j *Json[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	v, _ := (*datatypes.Json[T])(j).Value()
	return clause.Expr{
		SQL:  "?",
		Vars: []any{v},
	}
}

func (j *Json[T]) Value() (driver.Value, error) {
	// Scan a value into struct from database driver
	return (*datatypes.Json[T])(j).Value()
}

func (j *Json[T]) Scan(v any) error {
	// Scan a value into struct from database driver
	return (*datatypes.Json[T])(j).Scan(v)
}
