package usecase

import (
    "LayeredArchitecture/domain"
    "LayeredArchitecture/domain/repository"
    "gorm.io/gorm"
)

type ReportUsecase interface {
    SearchReport(req domain.SearchRequest, db *gorm.DB, userID uint) (*[]domain.Report, error)
    CreateReport(report *domain.Report, db *gorm.DB) error
    UpdateReport(report *domain.Report, db *gorm.DB) error
    DeleteReport(report *domain.Report, db *gorm.DB) error
    GetReport(id string, db *gorm.DB) (*domain.Report, error)
}

type reportUsecase struct {
    reportRepo repository.ReportRepository
}

func NewReportUsecase(rr repository.ReportRepository) ReportUsecase {
    return &reportUsecase{
        reportRepo: rr,
    }
}

func (r *reportUsecase) SearchReport(req domain.SearchRequest, db *gorm.DB, userID uint) (*[]domain.Report, error) {
    return r.reportRepo.SearchReport(req, db, userID)
}

func (r *reportUsecase) CreateReport(report *domain.Report, db *gorm.DB) error {
    return r.reportRepo.CreateReport(report, db)
}

func (r *reportUsecase) UpdateReport(report *domain.Report, db *gorm.DB) error {
    return r.reportRepo.UpdateReport(report, db)
}

func (r *reportUsecase) DeleteReport(report *domain.Report, db *gorm.DB) error {
    return r.reportRepo.DeleteReport(report, db)
}

func (r *reportUsecase) GetReport(id string, db *gorm.DB) (*domain.Report, error) {
    return r.reportRepo.GetReport(id, db)
}
