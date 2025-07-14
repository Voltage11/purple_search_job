package main

import (
	"lesson/config"
	"lesson/internal/home"
	"lesson/internal/users"
	"lesson/internal/vacancy"
	"lesson/pkg/database"
	"lesson/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	config.Init()
	dbConf := config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()

	customLogger := logger.NewLogger(logConfig)
	_ = dbConf

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) //middleware, при падении fiber не упадет приложение
	app.Static("/public", "./public")

	dbPool := database.CreateDbPool(dbConfig, customLogger)
	defer dbPool.Close()

	// Repositories
	vacancyRepo := vacancy.NewVacancyRepository(dbPool, customLogger)
	userRepo := users.NewUserRepository(dbPool, customLogger)


	// Handler
	home.NewHadnler(app, customLogger, vacancyRepo)
	vacancy.NewHadnler(app, customLogger, vacancyRepo)
	users.NewHadnler(app, customLogger, userRepo)

	app.Listen(":3000")
}