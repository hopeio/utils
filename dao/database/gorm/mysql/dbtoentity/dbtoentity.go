/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package dbtoentity

import (
	"github.com/hopeio/utils/dao/database/mysql"
	dbi "github.com/hopeio/utils/dao/database/toentity"
	"gorm.io/gorm"
)

func MysqlConvert(db *gorm.DB, filename string) {
	mysqlgen := mysqlgen{db: db}
	dbi.Convert(&mysqlgen, filename)
}

func MysqlConvertByTable(db *gorm.DB, tableName string) {
	mysqlgen := mysqlgen{db: db}
	dbi.ConvertByTable(&mysqlgen, tableName)
}

type mysqlgen struct {
	db *gorm.DB
}

func (m *mysqlgen) Tables() []string {
	var tables []string
	m.db.Raw(`SHOW TABLES`).Scan(&tables)
	return tables
}

func (m *mysqlgen) Fields(tableName string) []*dbi.Field {
	var dbfields []*dbi.Field
	m.db.Raw(`SHOW FULL COLUMNS FROM ` + tableName).Scan(&dbfields)
	return dbfields
}

func (m *mysqlgen) TypeToGoTYpe(typ string) string {
	return mysql.MySqlTypeToGoType(typ)
}
