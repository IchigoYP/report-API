package repository

import (
    "LayeredArchitecture/domain"
    "gorm.io/gorm"
)

type UserRepository interface {
    SearchUser(req domain.SearchRequest, db *gorm.DB) (*[]domain.User, error)
    CreateUser(user *domain.User, db *gorm.DB) error
    UpdateUser(user *domain.User, db *gorm.DB) error
    DeleteUser(user *domain.User, db *gorm.DB) error
    GetUser(id string, db *gorm.DB) (*domain.User, error)
}
