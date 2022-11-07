package model

import (
	// "encoding/json"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}
