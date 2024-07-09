package interfaces

import (
	"LayeredArchitecture/config"
	"LayeredArchitecture/domain"
	"LayeredArchitecture/usecase"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

// ReportHandler interface 定義
type ReportHandler interface {
	GetReport(c *gin.Context)
	SearchReport(c *gin.Context)
	CreateReport(c *gin.Context)
	UpdateReport(c *gin.Context)
	DeleteReport(c *gin.Context)
}

// reportHandler 構造体定義
type reportHandler struct {
	reportUsecase usecase.ReportUsecase
}

// NewReportHandler 関数は ReportHandler インターフェースを満たす reportHandler 構造体のインスタンスを返す
func NewReportHandler(rU usecase.ReportUsecase) ReportHandler {
	return &reportHandler{
		reportUsecase: rU,
	}
}

// GetReport メソッドの実装
func (h *reportHandler) GetReport(c *gin.Context) {
	id := c.Param("id")
	report, err := h.reportUsecase.GetReport(id, config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

// SearchReport メソッドの実装
func (h *reportHandler) SearchReport(c *gin.Context) {
	var req domain.SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search parameters"})
		return
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	reports, err := h.reportUsecase.SearchReport(req, config.DB, uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

// CreateReport メソッドの実装
func (h *reportHandler) CreateReport(c *gin.Context) {
	var report domain.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if !((report.Style == "だ・である調" || report.Style == "です・ます調") && (report.Language == "日本語" || report.Language == "英語")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not correct Style or Language"})
		return
	}
	userID, _ := strconv.Atoi(c.Param("id"))
	report.UserID = uint(userID)
	var task domain.Task
	task.ID = uint(userID)
	config.DB.Save(&task)
	report.TaskID = uint(userID)
	if err := h.reportUsecase.CreateReport(&report, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, report)
}

// UpdateReport メソッドの実装
func (h *reportHandler) UpdateReport(c *gin.Context) {
	var report domain.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if !((report.Style == "だ・である調" || report.Style == "です・ます調") && (report.Language == "日本語" || report.Language == "英語")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not correct Style or Language"})
		return
	}
	if err := h.reportUsecase.UpdateReport(&report, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}

// DeleteReport メソッドの実装
func (h *reportHandler) DeleteReport(c *gin.Context) {
	var report domain.Report
	reportID := c.Param("report_id")
	if err := config.DB.First(&report, reportID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}
	if err := h.reportUsecase.DeleteReport(&report, config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report deleted"})
}
