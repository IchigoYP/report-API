package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SearchRequest struct {
	Name        string `json:"name"`
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	IsCompleted *bool  `json:"iscompleted"`
	Style       string `json:"style"`
	Language    string `json:"language"`
}

type User struct {
	gorm.Model
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:255"`
	Reports   []Report
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Report struct {
	gorm.Model
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

type Task struct {
	gorm.Model
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := "localhost"
	dbPort := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&User{}, &Report{}, &Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()

	r.GET("/users", searchUserHandler(db))
	r.POST("/users", createUserHandler(db))
	r.PUT("/users/:id", updateUserHandler(db))
	r.DELETE("/users/:id", deleteUserHandler(db))
	r.GET("/users/:id", getUserHandler(db))

	r.GET("/users/:id/reports/:report_id", getReportHandler(db))
	r.POST("/users/:id/reports", createReportHandler(db))
	r.GET("/users/:id/reports", searchReportHandler(db))
	r.PUT("/users/:id/reports/:report_id", updateReportHandler(db))
	r.DELETE("/users/:id/reports/:report_id", deleteReportHandler(db))

	r.Run(":8080")
}

func searchUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var users []User
		searchUsers(req, db, &users)
		c.JSON(http.StatusOK, users)
	}
}

func searchUsers(req SearchRequest, db *gorm.DB, users *[]User) {
	query := db.Preload("Reports")
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	query.Find(users)
}

func searchReportHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SearchRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, _ := strconv.Atoi(c.Param("id"))
		var reports []Report
		searchReport(req, db, &reports, uint(id))
		c.JSON(http.StatusOK, reports)
	}
}

func searchReport(req SearchRequest, db *gorm.DB, reports *[]Report, id uint) {
	query := db
	if req.Title != "" {
		query = query.Where("title = ?", req.Title)
	}
	if req.IsCompleted != nil {
		query = query.Where("is_completed = ?", *req.IsCompleted)
	}
	if req.Language != "" {
		query = query.Where("language = ?", req.Language)
	}
	if req.Style != "" {
		query = query.Where("style = ?", req.Style)
	}
	query = query.Where("user_id = ?", id)
	query.Find(reports)
}

func createUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&user)
		c.JSON(http.StatusOK, user)
	}
}

func updateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&user)
		c.JSON(http.StatusOK, user)
	}
}

func updateReportHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var report Report
		reportID := c.Param("report_id")
		if err := db.First(&report, reportID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
			return
		}
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !((report.Style == "だ・である調" || report.Style == "です・ます調") && (report.Language == "日本語" || report.Language == "英語")) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not correct Style or Language"})
			return
		}
		db.Save(&report)
		c.JSON(http.StatusOK, report)
	}
}

func deleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		db.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}

func deleteReportHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var report Report
		reportID := c.Param("report_id")
		if err := db.First(&report, reportID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
			return
		}
		db.Delete(&report)
		c.JSON(http.StatusOK, gin.H{"message": "Report deleted"})
	}
}

func createReportHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.Preload("Reports").First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		var report Report
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !((report.Style == "だ・である調" || report.Style == "です・ます調") && (report.Language == "日本語" || report.Language == "英語")) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not correct Style or Language"})
			return
		}
		report.UserID = user.ID
		var task Task
		taskID := user.ID
		task.ID = user.ID
		report.TaskID = taskID
		db.Save(&task)
		db.Save(&report)
		db.Save(&user)
		c.JSON(http.StatusOK, report)
	}
}

func getReportHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var report Report
		id := c.Param("report_id")
		if err := db.First(&report, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "report not found"})
			return
		}
		c.JSON(http.StatusOK, report)
	}
}

func getUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		id := c.Param("id")
		if err := db.Preload("Reports").First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
