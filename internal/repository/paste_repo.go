package repository

import "gorm.io/gorm"

type pasteRepo struct {
	DB *gorm.DB
}

func NewPasteRepo(DB *gorm.DB) PasteRepo {
	return &pasteRepo{
		DB: DB,
	}
}
