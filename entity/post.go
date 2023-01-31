package entity

import "gorm.io/gorm"

type Post struct {
	gorm.Model // ini sudah mencakup id, dan timestamps
	Title   string `gorm:"type:VARCHAR(100);NOT NULL"`
	Content string `gorm:"type:LONGTEXT;NOT NULL"`
}