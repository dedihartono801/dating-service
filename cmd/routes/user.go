package routes

import (
	"dating-service/cmd/middleware"
	handler "dating-service/internal/delivery"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, signupHandler handler.UserHandler) {
	app.Post("/register", signupHandler.Signup)
	app.Post("/login", signupHandler.Login)
	app.Get("/users", middleware.AuthUser, signupHandler.UserList)
	app.Post("/swipe", middleware.AuthUser, signupHandler.Swipe)
}

func TransactionRouter(app fiber.Router, transactionHandler handler.TransactionHandler) {
	app.Post("/transaction", middleware.AuthUser, transactionHandler.Purchase)
	app.Get("/payment-method", middleware.AuthUser, transactionHandler.PaymentMethodList)
	app.Get("/package", middleware.AuthUser, transactionHandler.PackageList)
}
