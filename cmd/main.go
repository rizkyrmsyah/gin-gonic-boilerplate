package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/cmd/database"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/config"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/repository/mysql"
	transportHTTP "github.com/rizkyrmsyah/gin-gonic-boilerplate/transport/http"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/usecase"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewConnection(config.DB)

	r := gin.New()
	r.Use(logger.SetLogger())

	defer db.Close()

	// call repository
	example := mysql.NewUserRepository(db)

	//call usecase
	exampleUsecase := usecase.NewExample(example)

	//call transport
	transportHTTP.HttpExample(r, exampleUsecase)

	// run the service
	r.Run(":" + strconv.Itoa(config.AppPort))
}
