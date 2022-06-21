package models

import "github.com/jinzhu/gorm"

type Mahasiswa struct {
	gorm.Model
	Name       string `gorm:"size:255;not null;" json:"name"`
	Profilepic string `gorm:"size:255;not null;" json:"profilepic"`
	Email      string `gorm:"size:255;not null;unique" json:"email"`
	Password   string `gorm:"size:255;not null;" json:"password"`
	Bio        string `gorm:"size:255;not null;" json:"bio"`
	Link       string `gorm:"size:255;not null;" json:"link"`
	Whatsapp   string `gorm:"size:255;not null;" json:"whatsapp"`
	UserType   string `gorm:"size:255;not null;" json:"user_type"`
	Semester   string `gorm:"size:255;not null;" json:"semester"`
	Nim        string `gorm:"size:255;not null;" json:"nim"`
	Status     string `gorm:"size:255;not null;" json:"status"`
	//ForeignKey
	UniversityID uint `gorm:"not null" json:"university_id"`
	FakultasID   uint `gorm:"not null" json:"fakultas_id"`
	ProdiID      uint `gorm:"not null" json:"prodi_id"`

	Posts []Post `gorm:"foreignkey:UserID" json:"posts"`
}
