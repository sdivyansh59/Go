package main

import (
	"crmProject/database"
	"fmt"
	"crmProject/lead"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes (app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLeads)
	app.Delete("/api/v1/lead/:id",lead.DeleteLeads)

	
}


func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3","leads.db")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead)
	fmt.Println("Database Migrated")
}

func main() {
	fmt.Println("fiber rest api")
	app := fiber.New()

	setupRoutes(app)

    app.Listen(3000)

	defer database.DBConn.Close()
}