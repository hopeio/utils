package sql

func EndCallbackSQL(typ string) string {
	return `UPDATE cursor SET prev = next, cursor = '' WHERE type = '` + typ + `'`
}
