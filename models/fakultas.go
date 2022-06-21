package models

import "github.com/jinzhu/gorm"

type Fakultas struct {
	gorm.Model
	NameFakultas string `gorm:"size:255;not null;" json:"name_fakultas"`
	//foreign key to table user
	UniversityID uint `gorm:"not null" json:"university_id"`
}
