package excel

import "gorm.io/gorm"

// TODO
func export[T any](db *gorm.DB, filename string) error {
	var ts []T
	if err := db.Find(&ts).Error; err != nil {
		return err
	}
	return nil
}
