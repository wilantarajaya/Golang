package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)


func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result := config.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: book,
	})
}

func DetailBookController(c echo.Context) error {
	var books []models.Book


	var id, _ = strconv.Atoi(c.Param("id"))

	result := config.DB.First(&books, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: books,
	})
}

func GetAllBook(c echo.Context) error {
	var books []models.Book
	
	result := config.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: books,
	})
}

func DeleteBook(c echo.Context) error {
	var books []models.Book
	var id, _ = strconv.Atoi(c.Param("id"))

	result := config.DB.Delete(&books, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal Menghapus Data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success Menghapus data", Status: true, Data: books,
	})
}



