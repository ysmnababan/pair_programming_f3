package router

import (
	"pair_programming/controller"
	"pair_programming/repository"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, db *mongo.Database) {

	repo := &repository.Repo{DB: db}
	controller := &controller.TransactionController{TransactionRepo: repo}

	e.GET("/transaction/:id", controller.GetTransactionByID)
	e.GET("/transactions", controller.GetAllTransaction)
	e.POST("/transaction", controller.CreateNewTransaction)
}
