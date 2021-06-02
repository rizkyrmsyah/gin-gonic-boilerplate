package main

import (
	"log"
	"strconv"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/cmd/database"
	"github.com/rizkyrmsyah/gin-gonic-boilerplate/config"
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
	//

	//call usecase
	//

	//call transport
	//

	// run the service
	r.Run(":" + strconv.Itoa(config.AppPort))
}
