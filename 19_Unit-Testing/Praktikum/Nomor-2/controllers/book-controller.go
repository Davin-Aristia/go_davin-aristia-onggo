package controllers

import (
	"net/http"
	"strconv"

	"test/mvc/models"
	"test/mvc/repositories"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	responseData, err := repositories.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all books",
		"books":   responseData,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	responseData, err := repositories.GetBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get book with id " + strconv.Itoa(id),
		"book":    responseData,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	errBind := c.Bind(&book)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"error":   errBind.Error(),
		})
	}

	responseData, err := repositories.InsertBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new book",
		"book":    responseData,
	})

}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// Pengecekan apakah book dengan id tersebut ada dalam database
	_, err := repositories.GetBookByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, map[string]any{"message": "no book found for id "+strconv.Itoa(id)})
		}
	}

	err = repositories.DeleteBookByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete book with id " + strconv.Itoa(id),
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	updateData := models.Book{}
	errBind := c.Bind(&updateData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
			"error":   errBind.Error(),
		})
	}

	// Pengecekan apakah book dengan id tersebut ada dalam database
	book, err := repositories.GetBookByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, map[string]any{"message": "no book found for id "+strconv.Itoa(id)})
		}
	}

	responseData, err := repositories.UpdateBookByID(id, book, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update book with id " + strconv.Itoa(id),
		"book":    responseData,
	})
}
