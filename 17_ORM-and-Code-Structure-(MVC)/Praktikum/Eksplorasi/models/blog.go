package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserId uint `json:"userId" form:"userId"`
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}