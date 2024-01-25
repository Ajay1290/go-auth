package routes

import (
	"github.com/Ajay1290/go-auth/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Post("/api/user/register", controllers.RegisterUser)
	app.Post("/api/user/login", controllers.LoginUser)
	app.Post("/api/user/logout", controllers.LogoutUser)
	app.Get("/api/user", controllers.GetUser)

}
