package repository

import (
    "LayeredArchitecture/domain"
    "gorm.io/gorm"
)

type ReportRepository interface {
    SearchReport(req domain.SearchRequest, db *gorm.DB, userID uint) (*[]domain.Report, error)
    CreateReport(report *domain.Report, db *gorm.DB) error
    UpdateReport(report *domain.Report, db *gorm.DB) error
    DeleteReport(report *domain.Report, db *gorm.DB) error
    GetReport(id string, db *gorm.DB) (*domain.Report, error)
}
