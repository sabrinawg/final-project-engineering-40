package models

import "github.com/jinzhu/gorm"

type Post_mhs struct {
	gorm.Model
	Image   string `gorm:"size:255;not null;" json:"image"`
	Caption string `gorm:"size:255;not null;" json:"caption"`

	//foreign key to table mahasiswa
	MhsID uint `gorm:"not null" json:"mhs_id"`
}

func (p *Post_mhs) SavePostMahasiswa() (*Post_mhs, error) {
	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
