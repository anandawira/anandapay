package main

import (
	"net/http"
	"time"

	"github.com/anandawira/anandapay/config"
	"github.com/anandawira/anandapay/pkg/user/handler"
	"github.com/anandawira/anandapay/pkg/user/middleware"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/anandawira/anandapay/pkg/user/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.Connect()

	userRepo := repo.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, time.Second*5)
	handler.AttachHandler(r, userUsecase)
	r.GET("/test-jwt", middleware.Authenticate, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Authenticated",
		})
	})

	r.Run(":1234")
}
