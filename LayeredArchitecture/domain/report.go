package domain

import "time"

type Report struct {
	ID          uint      `gorm:"primary_key"`
	Title       string    `gorm:"size:255"`
	IsCompleted bool      `gorm:"default:false"`
	UserID      uint
	Count       uint
	TaskID      uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Style       string    `gorm:"size:10;default:'です・ます調'"`
	Language    string    `gorm:"size:10;default:'日本語'"`
}
