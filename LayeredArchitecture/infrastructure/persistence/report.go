package persistence

import (
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
    "gorm.io/gorm"
)

type reportRepository struct{}

func NewReportRepository() repository.ReportRepository {
    return &reportRepository{}
}

func (r *reportRepository) SearchReport(req domain.SearchRequest, db *gorm.DB, userID uint) (*[]domain.Report, error) {
    var reports []domain.Report
    query := db.Where("user_id = ?", userID)

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

    if err := query.Find(&reports).Error; err != nil {
        return nil, err
    }
    return &reports, nil
}

func (r *reportRepository) CreateReport(report *domain.Report, db *gorm.DB) error {
    return db.Create(report).Error
}

func (r *reportRepository) UpdateReport(report *domain.Report, db *gorm.DB) error {
    return db.Save(report).Error
}

func (r *reportRepository) DeleteReport(report *domain.Report, db *gorm.DB) error {
    return db.Delete(report).Error
}

func (r *reportRepository) GetReport(id string, db *gorm.DB) (*domain.Report, error) {
    var report domain.Report
    if err := db.First(&report, id).Error; err != nil {
        return nil, err
    }
    return &report, nil
}
