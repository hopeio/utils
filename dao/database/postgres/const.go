package postgres

import "github.com/hopeio/utils/dao/database"

const (
	ZeroTimeUCT     = "0001-01-01 00:00:00"
	ZeroTimeUCTZone = ZeroTimeUCT + "+00:00:00"
	ZeroTimeCST     = "0001-01-01 08:05:43"
	ZeroTimeCSTZone = ZeroTimeCST + "+08:05:43"
)

const (
	NotDeletedUCT = database.ColumnDeletedAt + " = '" + ZeroTimeUCT + "'"
	NotDeletedCST = database.ColumnDeletedAt + " = '" + ZeroTimeCST + "'"
)

const (
	WithNotDeletedUCT = ` AND ` + NotDeletedUCT
	WithNotDeletedCST = ` AND ` + NotDeletedUCT
)
