package controllers

import (
	"fmt"
	"log"
	"net/http"
	"rest-api-tutorial/portal/api/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	// if DbDriver == "mysql" {
	// 	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	// 	server.DB, err = gorm.Open(DbDriver, DbUrl)
	// 	if err != nil {
	// 		fmt.Printf("Can't connect to %s database", DbDriver)
	// 		log.Fatal("This is error: ", err)
	// 	} else {
	// 		fmt.Printf("We are connected to %s database", DbDriver)
	// 	}
	// }

	if DbDriver == "postgres" {
		DbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(DbDriver, DbUrl)
		if err != nil {
			fmt.Printf("Can't connect to %s database", DbDriver)
			log.Fatal("This is error: ", err)
		} else {
			fmt.Printf("We are connected to %s database", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	origins := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	accessOk := handlers.AllowCredentials()
	// throug := handlers.OptionsPassthrough(true)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(origins, headersOk, methodsOk, accessOk)(server.Router)))
}
