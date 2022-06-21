package models

import "github.com/jinzhu/gorm"

type Prodi struct {
	gorm.Model
	NameProdi    string `gorm:"size:255;not null;" json:"name_prodi"`
	UniversityID uint   `gorm:"not null" json:"university_id"`
	FakultasID   uint   `gorm:"not null" json:"fakultas_id"`
}
