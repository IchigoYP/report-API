package interfaces

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/domain"
	"LayeredArchitecture/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

// UserHandler interface 定義
type UserHandler interface {
	GetUser(c *gin.Context)
	SearchUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// userHandler 構造体定義
type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler 関数は UserHandler インターフェースを満たす userHandler 構造体のインスタンスを返す
func NewUserHandler(uU usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: uU,
	}
}

// GetUser メソッドの実装
func (h *userHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userUsecase.GetUser(id, config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// SearchUser メソッドの実装
func (h *userHandler) SearchUser(c *gin.Context) {
	var req domain.SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search parameters"})
		return
	}

	users, err := h.userUsecase.SearchUser(req, config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser メソッドの実装
func (h *userHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.userUsecase.CreateUser(&user, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser メソッドの実装
func (h *userHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.userUsecase.UpdateUser(&user, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser メソッドの実装
func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	if err := config.DB.Where("user_id = ?", user.ID).Delete(&domain.Report{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}
	if err := h.userUsecase.DeleteUser(&user, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
