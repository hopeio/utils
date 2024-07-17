package postgres

import (
	"database/sql"
	dbi "github.com/hopeio/utils/dao/database"
)

func ExistsByFilterExpressions(db *sql.DB, tableName string, filters dbi.FilterExprs) (bool, error) {
	result := db.QueryRow(`SELECT EXISTS(SELECT * FROM ` + tableName + `WHERE ` + filters.Build() + ` LIMIT 1)`)
	if err := result.Err(); err != nil {
		return false, err
	}
	var exists bool
	err := result.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
