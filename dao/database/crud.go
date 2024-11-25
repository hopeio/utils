/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package database

const (
	existsByColumnSQL = `SELECT EXISTS(SELECT * FROM %s WHERE %s = ?` + WithNotDeleted + ` LIMIT 1)`
	existsSQL         = `SELECT EXISTS(SELECT * FROM %s WHERE %s  LIMIT 1)`
	deleteSQL         = `Update %s SET deleted_at = now() WHERE %s = ?` + WithNotDeleted
)

func ExistsSQL(tableName, column string, withDeletedAt bool) string {
	sql := `SELECT EXISTS(SELECT * FROM ` + tableName + ` WHERE ` + column + ` = ?`
	if withDeletedAt {
		sql += WithNotDeleted
	}
	sql += ` LIMIT 1)`
	return sql
}

func DeleteSQL(tableName, column string) string {
	return `Update ` + tableName + ` SET deleted_at = now() WHERE ` + column + ` = ?` + WithNotDeleted
}

func ExistsSQLByQuerySQL(qsql string) string {
	return `SELECT EXISTS(` + qsql + ` LIMIT 1)`
}

func ExistsSQLByFilterExprs(tableName string, filters FilterExprs) string {
	return `SELECT EXISTS(SELECT * FROM ` + tableName + ` WHERE ` + filters.Build() + ` LIMIT 1)`
}
