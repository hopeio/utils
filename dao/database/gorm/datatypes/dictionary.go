package datatypes

import "gorm.io/gorm"

type Dictionary struct {
	Key   string
	Value string
}

func GetValue(db *gorm.DB, key string) (string, error) {
	var value string
	err := db.Table(`dictionary`).Select(`value`).Where(`key=?`, key).Scan(&value).Error
	if err != nil {
		return "", err
	}
	return value, nil
}

func SetValue(db *gorm.DB, key, value string) error {
	return db.Table(`dictionary`).Where(`key=?`, key).UpdateColumn("value", value).Error
}
