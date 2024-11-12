package main

import (
	"log"

	"dating-service/cmd/routes"
	"dating-service/database"
	transactionRepository "dating-service/internal/app/repository/transaction"
	repository "dating-service/internal/app/repository/user"
	"dating-service/internal/app/usecase/transaction"
	"dating-service/internal/app/usecase/user"
	handler "dating-service/internal/delivery"
	"dating-service/pkg/config"
	"dating-service/pkg/identifier"
	"dating-service/pkg/validator"

	validatorv10 "github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {
	swagger.New(swagger.Config{
		Title:        "Swagger API",
		DeepLinking:  false,
		DocExpansion: "none",
	})

	envConfig := config.SetupEnvFile()

	db := database.InitPostgres(envConfig)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())

	repository := repository.NewUserRepository(db)
	registerService := user.NewUserService(repository, validator, identifier)
	registerHandler := handler.NewSignupHandler(registerService)

	transactionRepository := transactionRepository.NewTransactionRepository(db)
	transactionService := transaction.NewTransactionService(transactionRepository, validator, identifier)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := fiber.New()

	routes.SetupRoutes(app)
	routes.UserRouter(app, registerHandler)
	routes.TransactionRouter(app, transactionHandler)

	if err := app.Listen(":5004"); err != nil {
		log.Fatalf("listen: %s", err)
	}
}
