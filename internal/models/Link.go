package models

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Hash    string `gorm:"type:varchar(100);unique_index"`
	LongURL string `gorm:"type:text"`
}
