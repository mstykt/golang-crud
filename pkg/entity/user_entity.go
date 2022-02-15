package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uint64 `gorm:"primaryKey"`
	Name           string
	Surname        string
	City           string
	IdentityNumber string `gorm:"index;unique;not null"`
	CreatedDate    time.Time
	UpdatedDate    time.Time
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdatedDate", time.Now())

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedDate", time.Now())
	tx.Statement.SetColumn("UpdatedDate", time.Now())

	return nil
}
