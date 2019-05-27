package user_service

import (
	"database/sql"
	"ohflutter-server/models"
	"time"
)

type User struct {
	UserID   int
	Mobile   string
	Name     string
	Age      sql.NullInt64
	Birthday *time.Time
	City     string
}

func (u *User) Add() error {

	user := map[string]interface{}{
		"mobile":   u.Mobile,
		"name":     u.Name,
		"age":      u.Age,
		"birthday": u.Birthday,
		"city":     u.City,
	}

	if err := models.AddUser(user); err != nil {
		return err
	}

	return nil
}
