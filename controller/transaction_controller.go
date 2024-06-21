package controller

import (
	"log"
	"net/http"
	"pair_programming/models"
	"pair_programming/repository"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController struct {
	TransactionRepo repository.TransactionRepo
}

func (c *TransactionController) GetTransactionByID(e echo.Context) error {
	id := e.Param("id")

	res, err := c.TransactionRepo.GetTransactionID(id)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, res)
}

func (c *TransactionController) GetAllTransaction(e echo.Context) error {
	res, err := c.TransactionRepo.GetAllTransaction()
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	return e.JSON(http.StatusOK, res)
}

func (c *TransactionController) CreateNewTransaction(e echo.Context) error {
	var t models.Transaction
	err := e.Bind(&t)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	// validate
	if t.ProductID == primitive.NilObjectID || t.UserID == primitive.NilObjectID || t.Quantity <= 0 {
		return e.JSON(http.StatusBadRequest, "error or missing param")
	}

	res, err := c.TransactionRepo.CreateTransaction(&t)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	// hit 3rd party api for generating invoice
	// insert code here, output is invoice link
	invoiceLink := ""

	return e.JSON(http.StatusCreated, map[string]interface{}{"result": res, "invoice": invoiceLink})
}
