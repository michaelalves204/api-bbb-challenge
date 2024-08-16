package models

import (
	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model
	Candidate string `json:"candidate"`
	Weigth    int    `json:"weigth" gorm:"default:1"`
}
