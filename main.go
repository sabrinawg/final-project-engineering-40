package main

import (
	"github.com/rg-km/final-project-engineering-40/controllers"
	"github.com/rg-km/final-project-engineering-40/middlewares"
	"github.com/rg-km/final-project-engineering-40/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()
	//set up router group gin
	r := gin.Default()

	//AUTH
	public := r.Group("/")
	public.GET("/cek", controllers.Test)
	public.POST("/signup", controllers.Register) //signup university
	public.POST("/signin", controllers.Login)    //signin university

	//ADMIN
	universitas := r.Group("/admin")
	universitas.Use(middlewares.JwtAuthMiddleware())                     //middleware jwt
	universitas.GET("/user", controllers.CurrentUser)                    // get current user
	universitas.PUT("/user", controllers.UpdateUser)                     // update user
	universitas.POST("/user/mhs", controllers.CreateMahasiswa)           // users/univ create user mahasiswa
	universitas.DELETE("/user/mhs/:id", controllers.DeleteMahasiswa)     // users/univ delete user mahasiswa
	universitas.POST("/user/fakultas", controllers.CreateFakultas)       // users/univ create user fakultas
	universitas.PUT("/user/fakultas/:id", controllers.UpdateFakultas)    // users/univ update user fakultas
	universitas.DELETE("/user/fakultas/:id", controllers.DeleteFakultas) // users/univ delete user fakultas
	universitas.POST("/user/prodi", controllers.CreateProdi)             // users/univ create user prodi
	universitas.PUT("/user/prodi/:id", controllers.UpdateProdi)          // users/univ update user prodi
	universitas.DELETE("/user/prodi/:id", controllers.DeleteProdi)       // users/univ delete user prodi
	universitas.POST("/user/post", controllers.CreatePost)               // users/univ create user post

	//MAHASISWA
	user := r.Group("/mhs")
	user.GET("/", controllers.GetAllMahasiswa)          // get all mahasiswa
	user.POST("/signin", controllers.LoginMahasiswa)    //signin mahasiswa
	user.Use(middlewares.JwtAuthMiddleware())           //middleware jwt
	user.POST("/post", controllers.CreatePostMahasiswa) //users/mhs create post

	//FAKULTAS
	fakultas := r.Group("/fakultas")
	fakultas.GET("/", controllers.GetAllFakultas)                       //get all fakultas
	fakultas.GET("/:id", controllers.GetFakultasByID)                   //get faculty by query(id)
	fakultas.GET("/name/:name_fakultas", controllers.GetFakultasByName) //get faculty by query(name)

	//PRODI
	prodi := r.Group("/prodi")
	prodi.GET("/", controllers.GetAllProdi)                    //get all prodi
	prodi.GET("/:id", controllers.GetProdiByID)                //get prodi by query (id)
	prodi.GET("/name/:name_prodi", controllers.GetProdiByName) //get prodi by query (name)

	//POST UNIVERSITAS
	post := r.Group("/post")
	post.GET("/", controllers.GetAllPost)

	//POST MAHASISWA
	postmhs := r.Group("/postmhs")
	postmhs.GET("/", controllers.GetAllPostMahasiswa)

	//server
	r.Run(":8080")
}
