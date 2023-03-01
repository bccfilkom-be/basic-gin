package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(50); NOT NULL" json:"name"`
	Username string `gorm:"type:VARCHAR(50); NOT NULL;UNIQUE" json:"username"`
	Password string `gorm:"type:TEXT; NOT NULL" json:"-"`
	Posts    []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"posts"`
}
