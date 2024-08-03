// user.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"railway/controllers"
	"railway/middleware"
)

func SetupUserRoutes(router fiber.Router) {
	router.Get("/me", middleware.DeserializeUser, controllers.GetMe)

	router.Post("/add-book", middleware.DeserializeUser, controllers.AddBookToCategory)
	router.Get("/dashboard", middleware.DeserializeUser, controllers.Dashboard)
	router.Get("/view-books", middleware.DeserializeUser, controllers.ViewBooks)
	// router.Get("/view-books", middleware.DeserializeUse,controllers.ViewBooks)
}
