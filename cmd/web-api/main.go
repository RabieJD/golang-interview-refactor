package main

import (
	"github.com/gin-gonic/gin"
	"interview/pkg/config"
	"interview/pkg/controllers"
	"interview/pkg/infra/database/cart/sql"
	"interview/pkg/infra/database/connection"
	"interview/pkg/infra/database/price/cache"
	"interview/pkg/service/calculator"
	"net/http"
)

func main() {
	// load config
	conf := config.New()

	// get database connection
	db, err := connection.GetDBConnection(&conf.DatabaseConnection)
	if err != nil {
		panic(err)
	}

	// create repos
	cartRepo, cartMigration := sql.NewRepo(db)
	priceRepo := cache.NewRepo()

	// migrate
	dbMigration := &connection.DBMigration{}
	dbMigration.Add(cartMigration)
	if err := dbMigration.MigrateDBs(); err != nil {
		panic(err)
	}

	// create cartService
	cartService := calculator.NewCartService(cartRepo, priceRepo)

	// create controllers
	taxController := controllers.NewTaxController(cartService)

	// add routes
	ginEngine := gin.Default()
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)

	// run server
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
