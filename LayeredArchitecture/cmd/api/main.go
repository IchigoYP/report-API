package main

import (
	"LayeredArchitecture/infrastructure/persistence"
	"LayeredArchitecture/interfaces/handler"
	"LayeredArchitecture/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	reportRepo := persistence.NewReportRepository()
	reportUsecase := usecase.NewReportUsecase(reportRepo)
	reportHandler := interfaces.NewReportHandler(reportUsecase)

	userRepo := persistence.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := interfaces.NewUserHandler(userUsecase)

	r := gin.Default()

	r.GET("/users/:id", userHandler.GetUser)
	r.GET("/users", userHandler.SearchUser)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	r.GET("/users/:id/reports/:report_id", reportHandler.GetReport)
	r.POST("/users/:id/reports", reportHandler.CreateReport)
	r.GET("/users/:id/reports", reportHandler.SearchReport)
	r.PUT("/users/:id/reports/:report_id", reportHandler.UpdateReport)
	r.DELETE("/users/:id/reports/:report_id", reportHandler.DeleteReport)

	r.Run(":8080")
}
