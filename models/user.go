package models

import (
	"errors"
	"html"
	"strings"

	"github.com/rg-km/final-project-engineering-40/utils/token"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Posts          []Post `gorm:"foreignkey:UserID" json:"posts"`
	NameUniversity string `gorm:"size:255;not null;" json:"name_university"`
	Profilepic     string `gorm:"null" json:"profilepic"`
	Email          string `gorm:"size:255;not null;unique" json:"email"`
	Password       string `gorm:"size:255;not null;" json:"password"`
	Bio            string `gorm:"null" json:"bio"`
	Link           string `gorm:"null" json:"link"`
	Whatsapp       string `gorm:"null" json:"whatsapp"`
	UserType       string `gorm:"null" json:"user_type"`

	//Register
	NameRektor string `gorm:"size:255;not null;" json:"name_rektor"`
	KtpRektor  string `gorm:"size:255;not null;" json:"ktp_rektor"`
	Isverified bool   `gorm:"default:false" json:"isverified"`
	Alamat     string `gorm:"null" json:"alamat"`
}

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}

func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *Mahasiswa) SaveMahasiswa() (*Mahasiswa, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Mahasiswa{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in email
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}

//get all mahasiswa
func GetAll() ([]Mahasiswa, error) {

	var mahasiswa []Mahasiswa

	err := DB.Find(&mahasiswa).Error
	if err != nil {
		return nil, err
	}
	return mahasiswa, nil
}
