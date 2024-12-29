/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package clause

import (
	dbi "github.com/hopeio/utils/dao/database/sql"
	"gorm.io/gorm/clause"
)

type ChainClause []clause.Interface

func (c ChainClause) ById(id int) ChainClause {
	if id != 0 {
		return c.ByIdNoCheck(id)
	}
	return c
}

func (c ChainClause) ByIdNoCheck(id any) ChainClause {
	return append(c, clause.Where{Exprs: []clause.Expression{clause.Eq{Column: dbi.ColumnId, Value: id}}})
}

func (c ChainClause) ByName(name string) ChainClause {
	if name != "" {
		return c.ByNameNoCheck(name)
	}
	return c
}

func (c ChainClause) ByNameNoCheck(name string) ChainClause {
	return append(c, clause.Where{Exprs: []clause.Expression{clause.Eq{Column: dbi.ColumnName, Value: name}}})
}
