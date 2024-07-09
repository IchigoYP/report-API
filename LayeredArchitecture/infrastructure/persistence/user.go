package persistence

import (
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
    "gorm.io/gorm"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
    return &userRepository{}
}

func (r *userRepository) SearchUser(req domain.SearchRequest, db *gorm.DB) (*[]domain.User, error){
    var users []domain.User
    query := db.Preload("Reports")
    if req.Name != "" {
        query = query.Where("name = ?", req.Name)
    }
    if req.ID != 0 {
        query = query.Where("id = ?", req.ID)
    }
    if err := query.Find(&users).Error; err != nil {
        return nil, err
    }
    return &users, nil
}

func (r *userRepository) CreateUser(user *domain.User, db *gorm.DB) error {
    return db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *domain.User, db *gorm.DB) error {
    return db.Save(user).Error
}

func (r *userRepository) DeleteUser(user *domain.User, db *gorm.DB) error {
    return db.Delete(user).Error
}

func (r *userRepository) GetUser(id string, db *gorm.DB) (*domain.User, error) {
    var user domain.User
    if err := db.Preload("Reports").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
