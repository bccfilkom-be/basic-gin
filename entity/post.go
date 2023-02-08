package entity

import "gorm.io/gorm"

// -> gorm.Model terdiri atas
/*
	type Model struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt DeletedAt `gorm:"index"`
	}
**/

type Post struct {
	gorm.Model           // ini sudah mencakup id, dan timestamps
	Title      string    `gorm:"type:VARCHAR(100);NOT NULL" json:"title"`
	Content    string    `gorm:"type:LONGTEXT;NOT NULL" json:"content"`
	UserID     uint      `json:"user_id"`
	Comments   []Comment `json:"comments"`
}
