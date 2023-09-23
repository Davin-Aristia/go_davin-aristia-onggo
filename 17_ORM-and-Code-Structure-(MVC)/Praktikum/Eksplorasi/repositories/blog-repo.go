package repositories

import (
	"mvc/eksplorasi/config"
	"mvc/eksplorasi/models"
)

func GetAllBlogs() ([]models.Blog, error) {
	var datablogs []models.Blog

	tx := config.DB.Find(&datablogs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datablogs, nil
}

func GetBlogByID(id int) (models.Blog, error) {
	var datablog models.Blog

	tx := config.DB.First(&datablog, id)
	if tx.Error != nil {
		return models.Blog{}, tx.Error
	}
	return datablog, nil
}

func InsertBlog(data models.Blog) (models.Blog, error) {
	tx := config.DB.Save(&data)
	if tx.Error != nil {
		return models.Blog{}, tx.Error
	}
	return data, nil
}

func DeleteBlogByID(id int) error {
	tx := config.DB.Delete(&models.Blog{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateBlogByID(id int, datablog, updateData models.Blog) (models.Blog, error) {
	tx := config.DB.Model(&datablog).Updates(map[string]interface{}{"title": updateData.Title, "content": updateData.Content})
	if tx.Error != nil {
		return models.Blog{}, tx.Error
	}
	return datablog, nil
}
