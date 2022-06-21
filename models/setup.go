package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", Dbdriver)
	}

	// auto migrate user table
	DB.AutoMigrate(&User{})
	// auto migrate fakultas table
	DB.AutoMigrate(&Fakultas{})
	// foreign key to table user
	DB.Model(&Fakultas{}).AddForeignKey("university_id", "users(id)", "CASCADE", "CASCADE")

	// auto migrate prodi table
	DB.AutoMigrate(&Prodi{})
	// foreign key to table users
	DB.Model(&Prodi{}).AddForeignKey("university_id", "users(id)", "CASCADE", "CASCADE")
	// foreign key to table fakultas
	DB.Model(&Prodi{}).AddForeignKey("fakultas_id", "fakultas(id)", "CASCADE", "CASCADE")

	// auto migrate mahasiswa table
	DB.AutoMigrate(&Mahasiswa{})
	// foreign key to table users
	DB.Model(&Mahasiswa{}).AddForeignKey("university_id", "users(id)", "CASCADE", "CASCADE")
	// foreign key to table fakultas
	DB.Model(&Mahasiswa{}).AddForeignKey("fakultas_id", "fakultas(id)", "CASCADE", "CASCADE")
	// foreign key to table prodi
	DB.Model(&Mahasiswa{}).AddForeignKey("prodi_id", "prodis(id)", "CASCADE", "CASCADE")

	// auto migrate post table
	DB.AutoMigrate(&Post{})
	// foreign key to table users
	DB.Model(&Post{}).AddForeignKey("univ_id", "users(id)", "CASCADE", "CASCADE")

	// auto migrate post_mhs table
	DB.AutoMigrate(&Post_mhs{})
	// foreign key to table users
	DB.Model(&Post_mhs{}).AddForeignKey("mhs_id", "mahasiswas(id)", "CASCADE", "CASCADE")
}
