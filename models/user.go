package models

import (
	"database/sql"
	"time"
)

type User struct {
	Model

	Mobile   string        `json:"mobile" gorm:"unique;not null"`
	Name     string        `json:"name"`
	Age      sql.NullInt64 `json:"age"`
	Birthday *time.Time
	City     string `json:"city"`
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Mobile:   data["mobile"].(string),
		Name:     data["name"].(string),
		Age:      data["age"].(sql.NullInt64),
		Birthday: data["birthday"].(*time.Time),
		City:     data["city"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
