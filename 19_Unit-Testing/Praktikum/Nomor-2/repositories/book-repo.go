package repositories

import (
	"test/mvc/config"
	"test/mvc/models"
)

func GetAllBooks() ([]models.Book, error) {
	var databooks []models.Book

	tx := config.DB.Find(&databooks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return databooks, nil
}

func GetBookByID(id int) (models.Book, error) {
	var databook models.Book

	tx := config.DB.First(&databook, id)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return databook, nil
}

func InsertBook(data models.Book) (models.Book, error) {
	tx := config.DB.Save(&data)
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return data, nil
}

func DeleteBookByID(id int) error {
	tx := config.DB.Delete(&models.Book{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateBookByID(id int, databook, updateData models.Book) (models.Book, error) {
	tx := config.DB.Model(&databook).Updates(map[string]interface{}{"title": updateData.Title, "writer": updateData.Writer, "publisher": updateData.Publisher})
	if tx.Error != nil {
		return models.Book{}, tx.Error
	}
	return databook, nil
}
