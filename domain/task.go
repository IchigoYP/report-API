package domain


type Task struct {
	ID        uint      
	Name      string    `gorm:"size:255"`
}
