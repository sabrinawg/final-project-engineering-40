package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Image   string `gorm:"size:255;not null;" json:"image"`
	Caption string `gorm:"size:255;not null;" json:"caption"`

	//foreign key to table user
	UnivID uint `gorm:"not null" json:"univ_id"`
}

func (p *Post) SavePost() (*Post, error) {
	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}
