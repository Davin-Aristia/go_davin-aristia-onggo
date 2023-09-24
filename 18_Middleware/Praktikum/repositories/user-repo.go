package repositories

import (
	"middleware/config"
	"middleware/models"
	"middleware/middlewares"
)

func CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User
	// select * from users where email = `email` and password = `password`
	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return models.User{}, "", tx.Error
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(data.ID))
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func GetAllUsers() ([]models.User, error) {
	var datausers []models.User

	tx := config.DB.Find(&datausers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datausers, nil
}

func GetUserByID(id int) (models.User, error) {
	var datauser models.User

	tx := config.DB.First(&datauser, id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return datauser, nil
}

func InsertUser(data models.User) (models.User, error) {
	tx := config.DB.Save(&data)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return data, nil
}

func DeleteUserByID(id int) error {
	tx := config.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateUserByID(id int, datauser, updateData models.User) (models.User, error) {
	tx := config.DB.Model(&datauser).Updates(map[string]interface{}{"name": updateData.Name, "email": updateData.Email, "password": updateData.Password})
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return datauser, nil
}
