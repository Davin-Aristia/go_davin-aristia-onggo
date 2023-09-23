package controllers

import (
	"net/http"
	"strconv"

	"mvc/eksplorasi/models"
	"mvc/eksplorasi/repositories"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	responseData, err := repositories.GetAllBlogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get all blogs",
		"blogs":   responseData,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	responseData, err := repositories.GetBlogByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error get data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get blog with id " + strconv.Itoa(id),
		"blog":    responseData,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	errBind := c.Bind(&blog)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind " + errBind.Error(),
			"error":   errBind,
		})
	}

	responseData, err := repositories.InsertBlog(blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success create new blog",
		"blog":    responseData,
	})

}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	// Pengecekan apakah blog dengan id tersebut ada dalam database
	_, err := repositories.GetBlogByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "no blog found for id "+strconv.Itoa(id))
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = repositories.DeleteBlogByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error delete data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete blog with id " + strconv.Itoa(id),
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	updateData := models.Blog{}
	c.Bind(&updateData)

	// Pengecekan apakah blog dengan id tersebut ada dalam database
	blog, err := repositories.GetBlogByID(id)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "no blog found for id "+strconv.Itoa(id))
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	responseData, err := repositories.UpdateBlogByID(id, blog, updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error update data",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success update blog with id " + strconv.Itoa(id),
		"blog":    responseData,
	})
}
