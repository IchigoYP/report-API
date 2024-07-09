package domain

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:255"`
	Reports   []Report
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
