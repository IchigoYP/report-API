package usecase

import (
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
    "gorm.io/gorm"
)

type UserUsecase interface {
    GetUser(id string, db *gorm.DB) (*domain.User, error)
    SearchUser(req domain.SearchRequest, db *gorm.DB) (*[]domain.User, error)
    CreateUser(user *domain.User, db *gorm.DB) error
    UpdateUser(user *domain.User, db *gorm.DB) error
    DeleteUser(user *domain.User, db *gorm.DB) error
}

type userUsecase struct {
    userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
    return &userUsecase{
        userRepo: ur,
    }
}

func (u *userUsecase) GetUser(id string, db *gorm.DB) (*domain.User, error) {
    return u.userRepo.GetUser(id, db)
}

func (u *userUsecase) SearchUser(req domain.SearchRequest, db *gorm.DB) (*[]domain.User, error) {
    return u.userRepo.SearchUser(req, db)
}

func (u *userUsecase) CreateUser(user *domain.User, db *gorm.DB) error {
    return u.userRepo.CreateUser(user, db)
}

func (u *userUsecase) UpdateUser(user *domain.User, db *gorm.DB) error {
    return u.userRepo.UpdateUser(user, db)
}

func (u *userUsecase) DeleteUser(user *domain.User, db *gorm.DB) error {
    return u.userRepo.DeleteUser(user, db)
}
